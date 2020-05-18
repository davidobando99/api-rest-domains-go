package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidobando99/APIRestWithGo/controller"
	"github.com/davidobando99/APIRestWithGo/model"
	"github.com/gorilla/mux"
)

var domain1 model.DomainApi

const HOST_NAME = "truora.com"

func main() {

	//fmt.Println(controller.GetPreviousGrade(domain1.Servers, controller.DomainList[0].SslGrade, controller.DomainList[0].LastTime))

	router := mux.NewRouter()
	router.HandleFunc("/domains", GetDomainsEndpoint).Methods("GET")
	router.HandleFunc("/domains/{id}", GetDomainEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getInfoJSON(host string) {
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

	domain1 = model.DomainApi{}
	_ = json.NewDecoder(response.Body).Decode(&domain1)
	//proving generate ssl grade

	domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "C"})
	/*
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "D"})
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "c"})
	*/
	fmt.Println(domain1.Servers)

	for _, server := range domain1.Servers {
		country, owner := model.WhoIsServer(server)
		fmt.Println(country)
		fmt.Println(owner)
	}
}

func GetDomainsEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(controller.DomainList)
}

func GetDomainEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	getInfoJSON(params["id"])
	controller.CreateDomainList(params["id"], model.GenerateSSLGrade(domain1.Servers), "B")
	json.NewEncoder(w).Encode(domain1.Servers)
}
