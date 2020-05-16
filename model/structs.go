package model

type Domain struct {
	Name             string    `json:"name"`
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
