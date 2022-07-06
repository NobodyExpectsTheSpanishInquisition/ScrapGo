package src

import (
	"fmt"
	"net/smtp"
)

type NotifierInterface interface {
	SendEmail(identity string, email Email) error
	getHost() string
	getPort() string
}

type Email struct {
	from    string
	to      string
	subject string
	body    string
}

type SmptNotifier struct {
	host     string
	port     string
	password string
}

func NewSmptNotifier(host string, port string, password string) *SmptNotifier {
	return &SmptNotifier{host: host, port: port, password: password}
}

func NewEmail(from string, to string, subject string, body string) Email {
	return Email{from: from, to: to, subject: subject, body: body}
}

func (n *SmptNotifier) getHost() string {
	return n.host
}

func (n *SmptNotifier) getPort() string {
	return n.port
}

func (n *SmptNotifier) SendEmail(identity string, email Email) error {
	var err error

	address := buildAddress(n)
	message := []byte(email.subject + email.body)
	auth := smtp.PlainAuth(identity, email.from, n.password, n.host)

	err = smtp.SendMail(address, auth, email.from, []string{email.to}, message)

	return err
}

func buildAddress(n NotifierInterface) string {
	return fmt.Sprintf("%s:%s", n.getHost(), n.getPort())
}
