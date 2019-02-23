package mailclient

type MailServer struct {
	Host    string
	Port    int
	User    string
	Pass    string
	UsedTLS bool //是否用TLS 发送邮件
}

const DEFAULTSENDCONTENTTYPE = "text/html; charset=UTF-8"
