package main

func main() {

	/*
		router := fasthttprouter.New()
		router.GET("/domains/", controller.GetDomainsEndpoint)
		router.GET("/domains/:host", controller.GetDomainEndpoint)
		fmt.Println("Server Listen at Port 8000")
		log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
	*/

	domainSQL.CreateTable()
	domainSQL.InsertDomain("truora.com", "A", "A")
	//fmt.Println()

}
