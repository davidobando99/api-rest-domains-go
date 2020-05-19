package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/davidobando99/APIRestWithGo/controller"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()
	router.GET("/domains/", controller.GetDomainsEndpoint)
	router.GET("/domains/:host", controller.GetDomainEndpoint)
	fmt.Println("Server Listen at Port 8000")
	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))

}
