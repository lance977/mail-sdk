package mail

// Attachment 邮件附件.
type Attachment struct {
	name string
}

// SetAttachment 设置附件名.
func (a *Attachment) SetAttachment(name string) *Attachment {
	return &Attachment{
		name: name,
	}
}

func (a *Attachment) Name() string {
	return a.name
}
