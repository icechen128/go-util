package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

// SendMail send mail of subject and body to args to
// to is ';' split email address
// subject is email title
// body is HTML string
func SendMail(to, subject, body string) error {
	user := "broadcast@blackvine.cn"
	password := "Blackvine123"
	host := "smtp.exmail.qq.com:587"
	err := sendMail(user, password, host, to, subject, body)
	if err != nil {
		return err
	}
	return nil
}

func sendMail(user, password, host, to, subject, body string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	fmt.Print("Send To Mail \n" + string(msg))
	err := smtp.SendMail(host, auth, user, strings.Split(to, ";"), msg)
	return err
}
