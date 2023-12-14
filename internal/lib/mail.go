package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
)

func SendCode(host, username, password, email, code string) error {
	t, err := template.ParseFiles("../internal/lib/templates/email-verify.html")

	if err != nil {
		return err
	}

	var body bytes.Buffer

	if err := t.Execute(&body, map[string]any{"ConfirmationURL": code}); err != nil {
		return err
	}

	res := make(chan error)
	message := []byte(
		"Subject: Код подтверждения\r\n" +
			"From: " + username + "\r\n" +
			"To: " + email + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" + body.String())

	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	go func() {
		res <- smtp.SendMail(host, auth, username, []string{email}, message)

	}()
	return <-res
}
