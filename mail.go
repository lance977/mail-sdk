package mail

import "context"

type Mailer interface {
	Check() error
	Send(ctx context.Context, msg *Message) error
}
