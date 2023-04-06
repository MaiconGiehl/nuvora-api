package usecase

type loginAsCompany struct {
	Email string
	Password string
}

func With(
	email string,
	password string,
	) *loginAsCompany {
	return &loginAsCompany{
		Email: email,
		Password: password,
	}
}