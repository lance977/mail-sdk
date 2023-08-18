package mail

import (
	"fmt"
	"strings"
)

type Address struct {
	name  string
	email string
}

func (a *Address) Name() string {
	return a.name
}

func (a *Address) Email() string {
	return a.email
}

func (a *Address) Address() string {
	return fmt.Sprintf("%s <%s>", a.name, a.email)
}

func (a *Address) SetAddress(name, email string) *Address {
	if name == "" {
		ms := strings.Split(email, "@")
		if len(ms) != 0 {
			name = ms[0]
		}
	}
	return &Address{
		name:  name,
		email: email,
	}
}
