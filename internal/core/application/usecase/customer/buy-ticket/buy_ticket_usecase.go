package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type BuyTicketUseCaseInterface interface {
	Execute(command *buyTicketCommand) (*[]dto.TicketOutputDTO, error)
}