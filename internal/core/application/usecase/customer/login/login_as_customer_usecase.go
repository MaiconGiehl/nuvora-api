package usecase

import entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"

type LoginAsCustomerUseCaseInterface interface {
	Execute(command *loginAsCustomerCommand) (*entity.Customer, error)
}