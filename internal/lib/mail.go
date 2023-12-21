package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
	"yata-email/config"
)

func SendCode(cfg config.SMTP, email, emailType, code string) error {
	templates := map[string]string{
		cfg.PasswordEmailType:     "../internal/lib/templates/password-verify.html",
		cfg.EmailConfirmationType: "../internal/lib/templates/email-verify.html",
	}

	t, err := template.ParseFiles(templates[emailType])

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
			"From: " + cfg.Username + "\r\n" +
			"To: " + email + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" + body.String())

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, "smtp.gmail.com")

	go func() {
		res <- smtp.SendMail(cfg.Host, auth, cfg.Username, []string{email}, message)

	}()
	return <-res
}
