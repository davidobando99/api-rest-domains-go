package controller

import (
	"time"

	"github.com/davidobando99/APIRestWithGo/model"
)

type DomainDB struct {
	Host        string
	SslGrade    string
	PreviousSSL string
	LastTime    time.Time
}

var DomainList []DomainDB

func CreateDomainList(hostname string, sslgrade string, previousssl string) {
	DomainList = append(DomainList, DomainDB{hostname, sslgrade, previousssl, time.Now()})
}

func CreateDomainDB(hostname string, sslgrade string, previousssl string) {

}

//sum 1 to the time hour saved on the list or DB and compare with the current time
func GetPreviousGrade(servers []model.ServerApi, gradeSSL string, lastTime time.Time) (string, string) {
	currentSSL := gradeSSL
	var newSSL string
	now := time.Now()
	last := lastTime.Add(1 * time.Hour)
	if last.After(now) {
		newSSL = model.GenerateSSLGrade(servers)
		return currentSSL, newSSL
	} else {
		newSSL = currentSSL
		return currentSSL, newSSL
	}

}
