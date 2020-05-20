package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/davidobando99/APIRestWithGo/controller"
	"github.com/valyala/fasthttp"
)

func main() {

	//db := database.Connection()
	//controller.DataBase = db
	router := fasthttprouter.New()
	router.GET("/domains/", controller.GetDomainsEndpoint)
	router.GET("/domains/:host", controller.GetDomainEndpoint)
	fmt.Println("Server Listen at Port 8000")
	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))

	/*
		database.CreateTable(db)
		database.InsertDomain(db, "hla.com", "A", "A")
		database.InsertDomain(db, "b.com", "A", "A")
		database.GetDomains(db)
		//fmt.Println()
	*/
	//servers := database.GetDomains(db)
	//fmt.Println("misServers")
	//fmt.Println(servers)
	//database.UpdateDomain(db, "hla.com", "B", "B")
	//database.SearchDomain("truora.com")
	//domain := database.SearchDomain(db, "truor.com")
	//fmt.Println(domain.Host)

}
