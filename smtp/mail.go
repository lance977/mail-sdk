package smtp

import (
	"context"

	"gopkg.in/gomail.v2"

	"github.com/lance977/mail-sdk"
)

const (
	DefaultPort = 465

	TextPlain = "text/plain"
	TextHtml  = "text/html"
)

var _ mail.Mailer = (*Mailer)(nil)

type Mailer struct {
	host     string
	port     int
	account  string
	password string
}

// NewMailer new smtp mailer.
func NewMailer(host string, port int, account, password string) *Mailer {
	if port <= 0 {
		port = DefaultPort
	}
	return &Mailer{
		host:     host,
		port:     port,
		account:  account,
		password: password,
	}
}

// Check mailer struct.
func (m *Mailer) Check() error {
	if m.host == "" {
		return mail.ErrEmptyHost
	}
	if m.port <= 0 {
		return mail.ErrPort
	}
	if m.account == "" {
		return mail.ErrEmptyAccount
	}
	if m.password == "" {
		return mail.ErrEmptyPassword
	}
	return nil
}

// Send email.
func (m *Mailer) Send(ctx context.Context, msg *mail.Message) error {
	err := m.Check()
	if err != nil {
		return err
	}

	if msg == nil {
		return mail.ErrEmptyMessage
	}

	mailMsg := gomail.NewMessage()
	// set from
	if msg.From() == nil {
		return mail.ErrEmptyFrom
	} else if msg.From().Email() == "" {
		return mail.ErrEmptyEmail
	} else {
		mailMsg.SetHeader("From", msg.From().Address())
	}

	// set recipients
	if len(msg.Recipients()) == 0 {
		return mail.ErrEmptyRecipient
	} else {
		recipients := make([]string, 0, len(msg.Recipients()))
		for _, address := range msg.Recipients() {
			recipients = append(recipients, address.Address())
		}
		mailMsg.SetHeader("To", recipients...)
	}

	// set cc
	if len(msg.Cc()) != 0 {
		cc := make([]string, 0, len(msg.Cc()))
		for _, address := range msg.Cc() {
			cc = append(cc, address.Address())
		}
		mailMsg.SetHeader("Cc", cc...)
	}

	// set bcc
	if len(msg.Bcc()) != 0 {
		bcc := make([]string, 0, len(msg.Bcc()))
		for _, address := range msg.Bcc() {
			bcc = append(bcc, address.Address())
		}
		mailMsg.SetHeader("Bcc", bcc...)
	}

	// set subject
	if msg.Subject() == "" {
		return mail.ErrEmptySubject
	} else {
		mailMsg.SetHeader("Subject", msg.Subject())
	}

	// body
	if msg.TextBody() != "" {
		mailMsg.SetBody(TextPlain, msg.TextBody())
	}
	if msg.HtmlBody() != "" {
		mailMsg.SetBody(TextHtml, msg.HtmlBody())
	}

	// attachments
	if len(msg.Attachments()) != 0 {
		for _, attachment := range msg.Attachments() {
			mailMsg.Attach(attachment.Name())
		}
	}

	// dial and send
	dial := gomail.NewDialer(m.host, m.port, m.account, m.password)

	return dial.DialAndSend(mailMsg)
}
