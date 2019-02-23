package mailclient

import (
	"fmt"
	"net/smtp"
)

func (s *MailServer)send(to []string, title, content, mailType string) error {

	mailAddr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	auth := smtp.PlainAuth(
		"",
		s.User,
		s.Pass,
		s.Host,
	)

	header := make(map[string]string)
	header["Subject"] = title
	header["Content-Type"] = mailType

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + content

	if s.UsedTLS {
		err := SendUsingTLS(
			mailAddr,
			auth,
			s.User,
			to,
			[]byte(message))
		return err
	}else {
		if err := smtp.SendMail(mailAddr, auth, s.User, to, []byte(message)); err!=nil{
			return err
		}
	}

	return nil
}

func (s *MailServer) SendHtml(to []string, title, content string) error {
	err := s.send(to, title, content, DEFAULTSENDCONTENTTYPE)
	return err
}



