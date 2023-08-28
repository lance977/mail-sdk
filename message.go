package mail

type Message struct {
	from        *Address      // 发送邮件地址
	to          []*Address    // 接收方邮件地址
	cc          []*Address    // 抄送
	bcc         []*Address    // 密送
	subject     string        // 主题
	textBody    string        // 文本内容
	htmlBody    string        // html内容
	attachments []*Attachment // 附件
}

// SetFrom 设置发件人地址.
func (msg *Message) SetFrom(addr *Address) {
	msg.from = addr
}

// From 获取发送方地址.
func (msg *Message) From() *Address {
	return msg.from
}

// SetRecipients 设置接收人地址.
func (msg *Message) SetRecipients(addresses []*Address) {
	msg.to = addresses
}

// Recipients 获取接收人地址.
func (msg *Message) Recipients() []*Address {
	return msg.to
}

// SetCc 设置抄送地址.
func (msg *Message) SetCc(addresses []*Address) {
	msg.cc = addresses
}

// Cc 获取抄送地址.
func (msg *Message) Cc() []*Address {
	return msg.cc
}

// SetBcc 设置密送地址.
func (msg *Message) SetBcc(addresses []*Address) {
	msg.bcc = addresses
}

// Bcc 获取密送地址.
func (msg *Message) Bcc() []*Address {
	return msg.bcc
}

// SetSubject 设置主题.
func (msg *Message) SetSubject(subject string) {
	msg.subject = subject
}

// Subject 获取主题.
func (msg *Message) Subject() string {
	return msg.subject
}

// TextBody 获取文本内容.
func (msg *Message) TextBody() string {
	return msg.textBody
}

// SetText 设置文本内容.
func (msg *Message) SetText(text string) {
	msg.textBody = text
}

// HtmlBody 获取html内容.
func (msg *Message) HtmlBody() string {
	return msg.htmlBody
}

// SetHtml 设置html邮件内容.
func (msg *Message) SetHtml(html string) {
	msg.htmlBody = html
}

// SetAttachments 设置附件.
func (msg *Message) SetAttachments(attachments []*Attachment) {
	msg.attachments = attachments
}

// Attachments 获取附件.
func (msg *Message) Attachments() []*Attachment {
	return msg.attachments
}
