package mail

type Attachment struct {
	name string
}

func (a *Attachment) SetAttachment(name string) *Attachment {
	return &Attachment{
		name: name,
	}
}

func (a *Attachment) Name() string {
	return a.name
}
