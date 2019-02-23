package mailclient

import "testing"

var sendMailTest = struct {
	Host    string
	Port    int
	User    string
	Pass    string
	UsedTLS bool
	To      []string
	Title   string
	Content string
}{
	Host:    "your host",
	Port:    25,
	UsedTLS: false,
	To:      []string{"mail list"},
	Title:   "mail title",
	Content: "mail content",
}

func TestSendHtmlMail(t *testing.T) {
	//t.Skip()
	s := NewServer(sendMailTest.Host, sendMailTest.Port, sendMailTest.UsedTLS)
	s = s.SetUser(sendMailTest.User, sendMailTest.Pass)

	err := s.SendHtml(sendMailTest.To, sendMailTest.Title, sendMailTest.Content)

	if err != nil {
		t.Errorf("send faild err:%v\n", err)
	}
}
