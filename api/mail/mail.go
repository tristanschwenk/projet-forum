package mail

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"text/template"

	"github.com/febytanzil/gomailer"
	"goyave.dev/goyave/v3/config"
)

// Initializer is a function meant to modify a connection settings
// at the global scope when it's created.
type Message gomailer.Message
type Attachment gomailer.Attachment

var (
	mailer gomailer.Client
	mu     sync.Mutex
)

func GetConnection() gomailer.Client {
	mu.Lock()
	defer mu.Unlock()
	if mailer == nil {
		mailer = newConnection()
	}
	return mailer
}

func Conn() gomailer.Client {
	return GetConnection()
}

func newConnection() gomailer.Client {
	mailer, _ := gomailer.NewClient(gomailer.Gomail, &gomailer.Config{
		Host:      config.GetString("mail.host"),
		Port:      config.GetInt("mail.port"),
		FromEmail: config.GetString("mail.fromEmail"),
		Username:  config.GetString("mail.username"),
		Password:  config.GetString("mail.password"),
	})

	return mailer
}

func (msg *Message) SendContext(ctx context.Context) error {
	mailer := GetConnection()
	return mailer.SendContext(ctx, (*gomailer.Message)(msg))
}

func (msg *Message) Send() error {
	mailer := GetConnection()
	return mailer.Send((*gomailer.Message)(msg))
}

func (msg *Message) SendAsync() error {
	mailer := GetConnection()
	return mailer.SendAsync((*gomailer.Message)(msg))
}

func (msg *Message) FromTemplate(templatePath string, data interface{}) error {

	if !strings.HasPrefix(templatePath, "./") {
		templatePath = GetTemplatePath(templatePath)
		fmt.Println(templatePath)
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	msg.Body = buf.String()
	return nil
}

func GetTemplatePath(name string) string {
	cur, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if !strings.HasSuffix(name, ".tmpl") {
		name = name + ".tmpl"
	}

	return path.Join(cur, "mail/templates", name)
}
