package usecase

import (
	"fmt"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)

type CreateCustomerAccountUseCase struct {
	CustomerRepository database.CustomerRepository
	PersonRepository database.PersonRepository
	AccountRepository database.AccountRepository
}

func NewCreateCustomerAccountUseCase(
	CustomerRepository database.CustomerRepository,
	PersonRepository database.PersonRepository,
	AccountRepository database.AccountRepository,
) *CreateCustomerAccountUseCase {
	return &CreateCustomerAccountUseCase{
		CustomerRepository: CustomerRepository,
		PersonRepository: PersonRepository,
		AccountRepository: AccountRepository,
	}
}

func (c *CreateCustomerAccountUseCase) Execute(input *dto.CustomerAccountInputDTO) error {
	customerEntity := entity.Customer{
		Name:   			input.Person.Customer.Name,
		Cpf:   				input.Person.Customer.Cpf,
		Phone:   			input.Person.Customer.Phone,
		BirthDate: 		input.Person.Customer.BirthDate,
		CompanyID:   	input.Person.Customer.CompanyID,
	}

	err := c.CustomerRepository.Save(&customerEntity)
	if err != nil {
		return err
	}

	newCustomerId, err := c.CustomerRepository.GetLastId()
	if err != nil {
		return err
	}

	personEntity := entity.Person{
  	Cep:							input.Person.Cep,
		PermissionLevel:	input.Person.PermissionLevel,
		CustomerID:				int(newCustomerId),
	}

	err = c.PersonRepository.SaveCustomerPerson(&personEntity)
	if err != nil {
		return err
	}

	newPersonId, err := c.PersonRepository.GetLastId()
	if err != nil {
		return err
	}

	accountEntity := entity.Account{
		Username: 			input.Username,
		Email: 					input.Email,
		Password: 			input.Password,
		PersonID: 			int(newPersonId),
		DailyTickets:   int(input.DailyTickets.Int16),
		TicketsLeft:    int(input.DailyTickets.Int16),
	}

	fmt.Print(accountEntity.TicketsLeft)

	err = c.AccountRepository.SaveCustomerAccount(&accountEntity)
	if err != nil {
		return err
	}

	return nil
}
