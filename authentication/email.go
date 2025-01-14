package authentication

func (svc *service) Login(email string, password string) (string, error) {

	err := svc.emailRepo.Send("Confirm email", "sourovsourov@gmail.com", "sourovsourov@gmail.com")
	if err != nil {
		return "", err
	}
	return "", nil
}
