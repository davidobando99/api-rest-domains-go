package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./model"
)

func main() {
	url := "https://api.ssllabs.com/api/v3/analyze?host="

	const HOST_NAME = "truora.com"
	url = url + "" + HOST_NAME
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

	domain1 := model.DomainApi{}
	_ = json.NewDecoder(response.Body).Decode(&domain1)
	//proving generate ssl grade
	/*
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "B"})
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "D"})
		domain1.Servers = append(domain1.Servers, model.ServerApi{"hola", "12.12", "c"})
	*/
	fmt.Println(domain1.Servers)

	for _, server := range domain1.Servers {
		country, owner := model.WhoIsServer(server)
		fmt.Println(country)
		fmt.Println(owner)
	}
	fmt.Println(model.GenerateSSLGrade(domain1.Servers))

}
