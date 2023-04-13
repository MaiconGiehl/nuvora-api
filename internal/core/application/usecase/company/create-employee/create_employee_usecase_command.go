package usecase

type createEmployeeCommand struct {
}

func With() *createEmployeeCommand {
	return &createEmployeeCommand{}
}
