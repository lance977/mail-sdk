package mail

import (
	"fmt"
	"strings"
)

// Address 邮件地址
type Address struct {
	name  string
	email string
}

// Name 获取名字.
func (a *Address) Name() string {
	return a.name
}

// Email 获取邮件地址.
func (a *Address) Email() string {
	return a.email
}

// Address 获取地址.
func (a *Address) Address() string {
	return fmt.Sprintf("%s <%s>", a.name, a.email)
}

// SetAddress 设置邮件地址, name可以为空, email不能为空.
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
