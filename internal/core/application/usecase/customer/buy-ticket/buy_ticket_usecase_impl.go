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

type BuyTicketUsecase struct {
	ctx context.Context
	logger logger.Logger
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository
}

func NewBuyTicketUsecase(
	ctx context.Context,
	logger logger.Logger,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *BuyTicketUsecase {
	return &BuyTicketUsecase{
		ctx: ctx,
		logger: logger,
		accountPGSQLRepository: accountPGSQLRepository,
		ticketPGSQLRepository: ticketPGSQLRepository,
		travelPGSQLRepository: travelPGSQLRepository,
	}	
}

func (u *BuyTicketUsecase) Execute(
	command *buyTicketCommand,
) error {

	err := u.validateInput(command)
	if err != nil {
		customErr := fmt.Sprintf("invalid path: %s", err.Error())
		err = errors.New(customErr)
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to buy ticket, %s", err.Error())
		return err
	}

	customerAccount, err := u.accountPGSQLRepository.GetAccountByID(command.AccountID)
	if err != nil {
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to find account, %s", err.Error())
		return errors.New("account not found")
	}

	if customerAccount.TicketsLeft.Int64 <= 0 {
		u.logger.Errorf("BuyTicketUsecase.Execute: Insufficient tickets")
		return errors.New("insuficient tickets")
	}

	err = u.accountPGSQLRepository.RemoveTicket(command.AccountID)
	if err != nil {
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to remove ticket from account, %s", err.Error())
		return err
	}

	err = u.ticketPGSQLRepository.CreateTicket(command.AccountID, command.TravelID)
	if err != nil {
		u.logger.Errorf("BuyTicketUsecase.Execute: Unable to buy ticket, %s", err.Error())
		return err
	}

	return err
}

func (u *BuyTicketUsecase) validateInput(input *buyTicketCommand) error {
	_, err := u.accountPGSQLRepository.GetAccountByID(input.AccountID)
	if err != nil {
		return errors.New("account not found")
	}

	_, err = u.travelPGSQLRepository.FindTravelByID(input.TravelID)
	if err != nil {
		return errors.New("travel not found")
	}
	
	return nil
}