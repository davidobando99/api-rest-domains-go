package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"./model"
	"github.com/likexian/whois-go"
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
	fmt.Println(domain1.Servers)
	for _, server := range domain1.Servers {
		country, owner := whoIsServer(server)
		fmt.Println(country)
		fmt.Println(owner)
	}

}

func whoIsServer(server model.ServerApi) (string, string) {
	ip := server.IP

	who, err := whois.Whois(ip)
	if err != nil {
		log.Fatal(err)
	}
	linesWho := (strings.Split(who, "\n"))
	var country string
	var owner string
	for i := 0; i < len(linesWho); i++ {
		if strings.Contains(linesWho[i], "Country") {
			country = linesWho[i]
			country = strings.Split(country, ":")[1]
			country = strings.TrimSpace(country)
		} else if strings.Contains(linesWho[i], "OrgName") {
			owner = linesWho[i]
			owner = strings.Split(owner, ":")[1]
			owner = strings.TrimSpace(owner)
		}
	}
	//fmt.Println(country)
	//fmt.Println(owner)
	return country, owner

}
