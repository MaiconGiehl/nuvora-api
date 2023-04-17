package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type BuyTicketUseCase struct {
	ctx                    context.Context
	logger                 logger.Logger
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	ticketPGSQLRepository  *ticket_entity.TicketPGSQLRepository
	travelPGSQLRepository  *travel_entity.TravelPGSQLRepository
}

func NewBuyTicketUseCase(
	ctx context.Context,
	logger logger.Logger,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *BuyTicketUseCase {
	return &BuyTicketUseCase{
		ctx:                    ctx,
		logger:                 logger,
		accountPGSQLRepository: accountPGSQLRepository,
		ticketPGSQLRepository:  ticketPGSQLRepository,
		travelPGSQLRepository:  travelPGSQLRepository,
	}
}

func (u *BuyTicketUseCase) Execute(
	command *buyTicketCommand,
) error {

	err := u.validateInput(command)
	if err != nil {
		customErr := fmt.Sprintf("invalid path: %s", err.Error())
		err = errors.New(customErr)
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to buy ticket, %s", err.Error())
		return err
	}

	customerAccount, err := u.accountPGSQLRepository.FindAccountByID(command.AccountID)
	if err != nil {
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to find account, %s", err.Error())
		return errors.New("account not found")
	}

	_, err = u.travelPGSQLRepository.FindTravelByID(command.TravelID)
	if err != nil {
		return errors.New("travel not found")
	}

	if customerAccount.TicketsLeft.Int64 <= 0 {
		u.logger.Errorf("BuyTicketUsecase.Execute: Insufficient tickets")
		return errors.New("insuficient tickets")
	}

	_ = u.accountPGSQLRepository.RemoveTicket(command.AccountID)

	_ = u.ticketPGSQLRepository.CreateTicket(command.AccountID, command.TravelID)

	return err
}

func (u *BuyTicketUseCase) validateInput(input *buyTicketCommand) error {
	_, err := u.accountPGSQLRepository.FindAccountByID(input.AccountID)
	if err != nil {
		return errors.New("account not found")
	}

	_, err = u.travelPGSQLRepository.FindTravelByID(input.TravelID)
	if err != nil {
		return errors.New("travel not found")
	}
	
	return nil
}
