package usecase

type loginAsCustomerCommand struct {
	Email string
	Password string
}

func With(
	email string,
	password string,
	) *loginAsCustomerCommand {
	return &loginAsCustomerCommand{
		Email: email,
		Password: password,
	}
}