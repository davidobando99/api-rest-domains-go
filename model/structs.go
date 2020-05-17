package model

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

type DomainApi struct {
	Host    string      `json:"host"`
	Servers []ServerApi `json:"endpoints"`
}
type ServerApi struct {
	Name string `json:"serverName"`
}
