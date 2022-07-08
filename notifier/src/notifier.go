package src

import (
	"fmt"
	"net/smtp"
)

type NotifierInterface interface {
	SendEmail(email Email) error
	getHost() string
	getPort() string
	getAddress() string
}

type SmptNotifier struct {
	host string
	port string
	auth smtp.Auth
}

type Email struct {
	from    string
	to      string
	subject string
	body    string
}

func newEmail(from string, to string, subject string, body string) Email {
	return Email{from: from, to: to, subject: subject, body: body}
}

func NewSmptNotifier(host string, port string, auth smtp.Auth) *SmptNotifier {
	return &SmptNotifier{host: host, port: port, auth: auth}
}

func (n *SmptNotifier) getPort() string {
	return n.port
}

func (n SmptNotifier) getAddress() string {
	return fmt.Sprintf("%s:%s", n.getHost(), n.getPort())
}

func (n *SmptNotifier) getHost() string {
	return n.host
}

func (e Email) getMessage() []byte {
	return []byte(e.subject + e.body)
}

func (n *SmptNotifier) SendEmail(email Email) error {
	return smtp.SendMail(n.getAddress(), n.auth, email.from, []string{email.to}, email.getMessage())
}

func CreateEmail(from string, to string, subject string, body string) Email {
	return newEmail(from, to, subject, body)
}
