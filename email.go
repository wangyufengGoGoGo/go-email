package go_email

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/smtp"
	"net/textproto"
	"strings"
	"time"
)

/**
 * @Author wyf
 * @Date 2021/8/10 9:27
 **/

type client struct {
	Option *ClientOption
}

type Client interface {
	SendEmail(request *Email) error
	SetHeader(email *Email) textproto.MIMEHeader
	HeaderToBytes(buff io.Writer, header textproto.MIMEHeader)
}

func (c *client) SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", c.Option.User, c.Option.Password, c.Option.Host)
	if len(email.To) <= 0 {
		return fmt.Errorf("accept people cannot be empty")
	}
	done := make(chan error, len(email.To))
	go func() {
		defer close(done)
		header := c.SetHeader(email)
		buf := bytes.NewBuffer(make([]byte, 0, 4096))
		c.HeaderToBytes(buf, header)
		io.WriteString(buf, email.Msg)

		err := smtp.SendMail(c.Option.ServerAddr, auth, c.Option.User, email.To, buf.Bytes())
		done <- err
	}()
	return <-done
}

func (c *client) SetHeader(email *Email) textproto.MIMEHeader {
	header := make(textproto.MIMEHeader)
	if _, ok := header["From"]; !ok {
		header.Set("From", c.Option.User)
	}
	if _, ok := header["To"]; !ok && len(email.To) > 0 {
		header.Set("To", strings.Join(email.To, ", "))
	}
	if _, ok := header["Subject"]; !ok && email.Subject != "" {
		header.Set("Subject", email.Subject)
	}
	if _, ok := header["Content-Type"]; !ok {
		header.Set("Content-Type", "text/plain;charset=UTF-8")
	}
	if _, ok := header["MIME-Version"]; !ok {
		header.Set("MIME-Version", "1.0")
	}
	if _, ok := header["Date"]; !ok {
		header.Set("Date", time.Now().Format(time.RFC1123Z))
	}

	return header
}

func (c *client) HeaderToBytes(buff io.Writer, header textproto.MIMEHeader) {
	for field, values := range header {
		for _, v := range values {
			io.WriteString(buff, field)
			io.WriteString(buff, ": ")
			switch {
			case field == "Content-Type":
				buff.Write([]byte(v))
			case field == "From" || field == "To":
				part := strings.Split(v, ",")
				buff.Write([]byte(strings.Join(part, ", ")))
			default:
				buff.Write([]byte(mime.QEncoding.Encode("UTF-8", v)))
			}
		}
		io.WriteString(buff, "\r\n")
	}
	io.WriteString(buff, "\r\n")
}

func NewClient(option *ClientOption) Client {
	return &client{Option: option}
}
