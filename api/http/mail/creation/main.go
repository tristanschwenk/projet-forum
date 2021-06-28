package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "noreply@ezyostudio.com")

	// Set E-Mail receivers
	m.SetHeader("To", "zao.soula@ezyostudio.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", "<h1> Hello world </h1> <p> Successfully created your password </p>")

	// Settings for SMTP server
	d := gomail.NewDialer("mail.gandi.net", 587, "tristan.schwenk@ezyostudio.com", "xuqqip-xapvE6-zunnuh")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
