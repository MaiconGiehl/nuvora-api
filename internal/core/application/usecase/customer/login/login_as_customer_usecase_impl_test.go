package usecase

import (
	"context"
	"testing"

	"github.com/maicongiehl/nuvora-api/configs/env"
	postgresdb "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	logrus_config "github.com/maicongiehl/nuvora-api/internal/infra/log/logrus"

	"github.com/stretchr/testify/suite"
)

var ctx = context.Background()

type LoginUseCaseImplTestSuite struct {
	suite.Suite
	ctx             context.Context
	logger          *logrus_config.LogrusLogger
	loginRepository *LoginAsCustomerUseCase
}

func (suite *LoginUseCaseImplTestSuite) SetupTest() {
	logger := logrus_config.NewLogrusLogger()

	suite.ctx = ctx
	suite.logger = logger
	env := env.LoadConfig("../../../../../../.env", logger)

	db := postgresdb.ConnectWithDB(
		logger,
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)

	customerPGSQLRepository := customer_entity.NewCustomerPGSQLRepository(ctx, db, logger)
	personPGSQLRepository := person_entity.NewPersonPGSQLRepository(ctx, db, logger)
	accountPGSQLRepository := account_entity.NewAccountPGSQLRepository(ctx, db, logger)
	useCaseRepository := NewLoginAsCustomerUseCase(
		ctx,
		logger,
		customerPGSQLRepository,
		personPGSQLRepository,
		accountPGSQLRepository,
	)

	suite.loginRepository = useCaseRepository

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(LoginUseCaseImplTestSuite))
}

func (s *LoginUseCaseImplTestSuite) TestLogin() {
	command := With(
		"usuario1@gmail.com",
		"usuario1",
	)
	_, err := s.loginRepository.Execute(command)
	s.Error(err)

	command = With(
		"empresa01@gmail.com",
		"empresa01",
	)
	_, err = s.loginRepository.Execute(command)
	s.Error(err)

	command = With(
		"usuario01@gmail.com",
		"usuario01",
	)
	_, err = s.loginRepository.Execute(command)
	s.NoError(err)
}
