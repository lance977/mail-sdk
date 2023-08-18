package mail

type Message struct {
	from        *Address
	to          []*Address
	cc          []*Address
	bcc         []*Address
	subject     string
	textBody    string
	htmlBody    string
	attachments []*Attachment
}

func (msg *Message) SetFrom(addr *Address) {
	msg.from = addr
}

func (msg *Message) From() *Address {
	return msg.from
}

func (msg *Message) SetRecipients(addresses []*Address) {
	msg.to = addresses
}

func (msg *Message) Recipients() []*Address {
	return msg.to
}

func (msg *Message) SetCc(addresses []*Address) {
	msg.cc = addresses
}

func (msg *Message) Cc() []*Address {
	return msg.cc
}

func (msg *Message) SetBcc(addresses []*Address) {
	msg.bcc = addresses
}

func (msg *Message) Bcc() []*Address {
	return msg.bcc
}

func (msg *Message) SetSubject(subject string) {
	msg.subject = subject
}

func (msg *Message) Subject() string {
	return msg.subject
}

func (msg *Message) TextBody() string {
	return msg.textBody
}

func (msg *Message) SetText(text string) {
	msg.textBody = text
}

func (msg *Message) HtmlBody() string {
	return msg.htmlBody
}

func (msg *Message) SetHtml(html string) {
	msg.htmlBody = html
}

func (msg *Message) SetAttachments(attachments []*Attachment) {
	msg.attachments = attachments
}

func (msg *Message) Attachments() []*Attachment {
	return msg.attachments
}
