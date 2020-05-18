package model

import (
	"log"
	"strings"

	"github.com/likexian/whois-go"
)

type Domain struct {
	HostName         string    `json:"host"`
	Servers          *[]Server `json:"servers"`
	ServersChanged   bool      `json:"servers_changed"`
	SslGrade         string    `json:"ssl_grade"`
	PreviousSslGrade string    `json:"previous_ssl_grade"`
	Logo             string    `json:"logo”:"`
	Title            string    `json:"title"`
	IsDown           bool      `json:"is_down"`
}

type Server struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Owner    string `json:"owner"`
	Country  string `json:"country"`
}

type DomainJson struct {
	Host    string       `json:"host"`
	Servers []ServerJson `json:"endpoints"`
	Errors  []ErrorJson  `json:"errors"`
}
type ServerJson struct {
	Name  string `json:"serverName"`
	IP    string `json:"ipAddress"`
	Grade string `json:"grade"`
}

type ErrorJson struct {
	Message string `json:"message"`
}

func WhoIsServer(server ServerJson) (string, string) {
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

func GenerateSSLGrade(servers []ServerJson) string {
	var minorGrade string
	availableGrades := []string{"A", "B", "C", "D", "E", "F"}
	if len(servers) == 1 {
		if existSSL(availableGrades, servers[0].Grade) {
			return servers[0].Grade
		}
	} else {
		minorGrade = servers[0].Grade
		for _, server := range servers[1:] {
			if existSSL(availableGrades, server.Grade) {
				if server.Grade > minorGrade {
					minorGrade = server.Grade
				}
			}
		}
	}
	return minorGrade
}

func existSSL(arr []string, find string) bool {
	for _, ssl := range arr {
		if ssl == find {
			return true
		}
	}
	return false
}
