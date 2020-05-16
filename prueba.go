package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type peopl struct {
	Number string `json:"message"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	//Create a request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")
	//Client do a request
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	//Take body responses
	//body, readErr := ioutil.ReadAll(res.Body)
	//if readErr != nil {
	//	log.Fatal(readErr)
	//}

	people1 := peopl{}
	//jsonErr := json.Unmarshal(body, &people1)
	_ = json.NewDecoder(res.Body).Decode(&people1)
	//if jsonErr != nil {
	//	log.Fatal(jsonErr)
	//}

	fmt.Println(people1.Number)
}
