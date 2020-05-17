package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidobando99/APIRestWithGo/model"
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
}
