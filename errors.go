package mail

import "github.com/pkg/errors"

var (
	ErrEmptyHost      = errors.New("host is empty")
	ErrPort           = errors.New("invalid port")
	ErrEmptyAccount   = errors.New("account is empty")
	ErrEmptyPassword  = errors.New("password is empty")
	ErrEmptyToken     = errors.New("token is empty")
	ErrEmptyEmail     = errors.New("email address is empty")
	ErrEmptyFrom      = errors.New("from is empty")
	ErrEmptyRecipient = errors.New("recipient is empty")
	ErrEmptyMessage   = errors.New("message is nil")
	ErrEmptySubject   = errors.New("subject is empty")
)
