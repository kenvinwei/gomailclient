package mailclient


func NewServer(host string,port int, usedTLS bool) *MailServer {
	return &MailServer{
		Host:host,
		Port:port,
		UsedTLS:usedTLS,
	}
}

func (s *MailServer)AddUser(user,pass string) *MailServer {
	if len(s.User) == 0 {
		s.User = user
	}

	if len(s.Pass) == 0 {
		s.Pass = pass
	}

	return s
}
