package main

import (
	"crypto/tls"
	"fmt"

	"github.com/Unknwon/goconfig"
	"gopkg.in/gomail.v2"
)

func SendMail(subject string, body string) {
	var session string = "mail"

	c, err := goconfig.LoadConfigFile("set.ini")
	if err != nil {
		panic(err)
	}

	from, _ := c.GetValue(session, "from")
	tolist := c.MustValueArray(session, "to", ",")
	host, _ := c.GetValue(session, "host")
	port, _ := c.Int(session, "port")
	username, _ := c.GetValue(session, "username")
	password, _ := c.GetValue(session, "password")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tolist...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {

		panic(err)
	}
}

func main() {

	SendMail("12345", "djjdei djei djeijdie jidjie ")
	fmt.Printf("fin\n")
}
