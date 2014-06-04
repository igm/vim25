package vim25

func (s *VimService) Login(this *SessionManager, login, pass string) error {
	loginResponse := new(LoginResponse)
	opLogin := Login{
		This:     this,
		Username: login,
		Password: pass,
	}
	if err := s.Invoke(opLogin, loginResponse); err != nil {
		return err
	}
	return nil
}
