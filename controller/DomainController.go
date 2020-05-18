package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidobando99/APIRestWithGo/controller"
	"github.com/davidobando99/APIRestWithGo/model"
	"github.com/valyala/fasthttp"
)

type DomainDB struct {
	Host        string
	SslGrade    string
	PreviousSSL string
	LastTime    time.Time
}

var DomainList []DomainDB

func CreateDomainList(hostname string, sslgrade string, previousssl string) {
	DomainList = append(DomainList, DomainDB{hostname, sslgrade, previousssl, time.Now()})
}

func CreateDomainDB(hostname string, sslgrade string, previousssl string) {

}

//sum 1 to the time hour saved on the list or DB and compare with the current time
func GetPreviousGrade(servers []model.ServerJson, gradeSSL string, lastTime time.Time) (string, string) {
	currentSSL := gradeSSL
	var newSSL string
	now := time.Now()
	last := lastTime.Add(1 * time.Hour)
	if last.After(now) {
		newSSL = model.GenerateSSLGrade(servers)
		return currentSSL, newSSL
	} else {
		newSSL = currentSSL
		return currentSSL, newSSL
	}

}
func getInfoJSON(host string) model.DomainJson {
	url := "https://api.ssllabs.com/api/v3/analyze?host="

	url = url + "" + host
	cliente := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
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
	//proving generate ssl grade

	domain.Servers = append(domain.Servers, model.ServerJson{"hola", "12.12", "C"})
	/*
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "D"})
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "c"})
	*/
	fmt.Println(domain.Servers)

	for _, server := range domain.Servers {
		country, owner := model.WhoIsServer(server)
		fmt.Println(country)
		fmt.Println(owner)
	}

	return domain
}
func GetDomainsEndpoint(ctx *fasthttp.RequestCtx) {
	doJSONWrite(ctx, fasthttp.StatusOK, controller.DomainList)
}
func GetDomainEndpoint(ctx *fasthttp.RequestCtx) {
	host := ctx.UserValue("host").(string)
	fmt.Println(host)

	doJSONWrite(ctx, fasthttp.StatusOK, DomainFromJsonApi(host))
}

func DomainFromJsonApi(host string) model.Domain {
	var domain model.Domain

	domainJson := getInfoJSON(host)
	isDown := len(domainJson.Errors) != 0
	domain.HostName = host
	domain.Servers = ServersFromJsonApi(domainJson.Servers)
	domain.SslGrade = model.GenerateSSLGrade(domainJson.Servers)
	domain.PreviousSslGrade = GetPreviousGrade(domainJson.Servers, domain.SslGrade)
	domain.IsDown = isDown
	controller.CreateDomainList(host, domain.SslGrade, domain.PreviousSslGrade)

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
