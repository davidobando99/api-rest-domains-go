package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()
	router.GET("/domains/", controller.GetDomainsEndpoint)
	router.GET("/domains/:host", controller.GetDomainEndpoint)
	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
