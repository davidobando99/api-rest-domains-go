package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/davidobando99/APIRestWithGo/database"

	"github.com/badoux/goscraper"
	"github.com/davidobando99/APIRestWithGo/model"
	"github.com/valyala/fasthttp"
)

var Consult model.ConsultedDomains
var DataBase *sql.DB

func GetPreviousGrade(servers []model.ServerJson, gradeSSL string, lastTime time.Time) (string, string) {
	previousSSL := gradeSSL
	currentSSL := model.GenerateSSLGrade(servers)

	if previousSSL == currentSSL {

		return previousSSL, previousSSL
	}

	return previousSSL, currentSSL

}

//sum 1 to the time hour saved on the list or DB and compare with the current time
func ServerHasChanged(previousSSL string, currentSSL string, lastTime time.Time) bool {
	now := time.Now()
	last := lastTime.Add(1 * time.Hour)
	if last.Before(now) && previousSSL != currentSSL {
		return true
	}
	return false

}
func getInfoJSON(host string) model.DomainJson {
	url := "https://api.ssllabs.com/api/v3/analyze?host="

	url = url + "" + host
	cliente := http.Client{
		Timeout: time.Second * 2,
	}
	//Create a request
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Client do a request
	response, getErr := cliente.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	domain := model.DomainJson{}
	_ = json.NewDecoder(response.Body).Decode(&domain)

	return domain
}
func GetDomainsEndpoint(ctx *fasthttp.RequestCtx) {
	GetDomainsFromDatabase(DataBase)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	doJSONWrite(ctx, fasthttp.StatusOK, Consult)
	Consult = model.ConsultedDomains{}
}
func GetDomainEndpoint(ctx *fasthttp.RequestCtx) {
	host := ctx.UserValue("host").(string)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	doJSONWrite(ctx, fasthttp.StatusOK, DomainFromJsonApi(DataBase, host))
}

func DomainFromJsonApi(db *sql.DB, host string) model.Domain {
	var domain model.Domain
	var sslGrade, previousGrade, title, logo string
	var lastTime time.Time
	var serversJson []model.ServerJson

	domainJson := getInfoJSON(host)
	fmt.Println(domainJson.Servers)
	if len(domainJson.Servers) == 0 {
		domainJson = getInfoJSON(host)
	}
	isDown := len(domainJson.Errors) != 0 || domainJson.Status == "ERROR"
	founded := database.SearchDomain(db, host)
	//SI ESTA CAIDO EL SERVER SU LISTA DE ENPOINTS SERA VACIO, SINO SERA LAS OBTENIDAS POR EL JSON
	if isDown {
		serversJson = []model.ServerJson{}
	} else {
		serversJson = domainJson.Servers
		logo, title = GetLogoAndTitle("http://www." + host)
	}
	if strings.Compare(founded.Host, "") == 0 {
		sslGrade = model.GenerateSSLGrade(serversJson)
		previousGrade = sslGrade
		database.InsertDomain(db, host, sslGrade, previousGrade)
		lastTime = time.Now()
	} else {
		currentGrade := founded.SslGrade
		lastTime = founded.LastTime
		if isDown {
			sslGrade = currentGrade
			previousGrade = founded.PreviousSSL
		} else {
			previousGrade, sslGrade = GetPreviousGrade(serversJson, currentGrade, lastTime)
		}
		database.UpdateDomain(db, host, sslGrade, previousGrade)

	}
	domain = CreateDomain(serversJson, host, sslGrade, previousGrade, title, logo, isDown, lastTime)

	return domain
}

func CreateDomain(serversJson []model.ServerJson, host string, sslGrade string, previousGrade string, title string, logo string, isDown bool, lastTime time.Time) model.Domain {
	var domain model.Domain
	var servers []model.Server

	if !isDown {
		servers = ServersFromJsonApi(serversJson)
	}
	domain.HostName = host
	domain.Servers = servers
	domain.SslGrade = sslGrade
	domain.PreviousSslGrade = previousGrade
	domain.Title = title
	domain.Logo = logo
	domain.IsDown = isDown
	domain.ServersChanged = ServerHasChanged(previousGrade, sslGrade, lastTime) //True si el ssl grade es distinto al que tenia el server una hora o mas antes
	return domain
}

func ServersFromJsonApi(serversJson []model.ServerJson) []model.Server {
	var servers []model.Server
	for _, server := range serversJson {
		country, owner := model.WhoIsServer(server)
		serverNew := model.Server{
			Address:  server.IP,
			SslGrade: server.Grade,
			Owner:    owner,
			Country:  country,
		}
		servers = append(servers, serverNew)
	}
	return servers
}
func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	var (
		strContentType     = []byte("Content-Type")
		strApplicationJSON = []byte("application/json")
	)
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func GetLogoAndTitle(url string) (string, string) {
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	return s.Preview.Icon, s.Preview.Title
}

func GetDomainsFromDatabase(db *sql.DB) {
	domains := database.GetDomains(db)
	for _, domain := range domains {
		//if !VerifyExistedDomain(Consult.Items, domain.Host) {
		item := model.Item{
			domain.Host,
		}
		Consult.Items = append(Consult.Items, item)

		//}

	}
}

func VerifyExistedDomain(domains []model.Item, host string) bool {

	for _, domain := range domains {
		if domain.HostName == host {
			return true
			break
		}
	}

	return false

}
