package usecase

import (
	"context"
	"testing"

	"github.com/maicongiehl/nuvora-api/configs/env"
	postgresdb "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
	logrus_config "github.com/maicongiehl/nuvora-api/internal/infra/log/logrus"

	"github.com/stretchr/testify/suite"
)

var ctx = context.Background()

type BuyTicketUseCaseImplTestSuite struct {
	suite.Suite
	ctx                 context.Context
	logger              *logrus_config.LogrusLogger
	buyTicketRepository *BuyTicketUseCase
}

func (suite *BuyTicketUseCaseImplTestSuite) SetupTest() {
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

	travelPGSQLRepository := travel_entity.NewTravelPGSQLRepository(ctx, db, logger)
	ticketPGSQLRepository := ticket_entity.NewTicketPGSQLRepository(ctx, db, logger)
	accountPGSQLRepository := account_entity.NewAccountPGSQLRepository(ctx, db, logger)

	useCaseRepository := NewBuyTicketUseCase(
		ctx,
		logger,
		accountPGSQLRepository,
		ticketPGSQLRepository,
		travelPGSQLRepository,
	)

	suite.buyTicketRepository = useCaseRepository

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(BuyTicketUseCaseImplTestSuite))
}

func (s *BuyTicketUseCaseImplTestSuite) TestBuyTicket() {
	command := With(1, 1)

	s.buyTicketRepository.accountPGSQLRepository.AddTicket(command.AccountID)

	s.NoError(s.buyTicketRepository.Execute(command))

	command = With(-1, 1)
	s.Error(s.buyTicketRepository.Execute(command))

	command = With(1, -1)
	s.Error(s.buyTicketRepository.Execute(command))
}
