package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type GetLastPurchasesUseCaseInterface interface {
	Execute(command *getLastPurchasesCommand) (*[]dto.TicketOutputDTO, error)
}