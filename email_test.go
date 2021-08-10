package go_email

import (
	"fmt"
	"testing"
)

/**
 * @Author wyf
 * @Date 2021/8/9 16:50
 **/

func TestSendEmail(t *testing.T) {
	option := &ClientOption{
		Host:       "smtp.163.com",
		ServerAddr: "smtp.163.com:25",
		User:       "******@163.com",
		Password:   "******",
	}
	c := NewClient(option)
	email := &Email{
		To:      []string{"******@163.com"},
		Subject: "hello email",
		Msg:     "Hello World , I am Golang",
	}
	errors := c.SendEmail(email)
	fmt.Println(errors)
}
