package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type GetPurchasesUseCaseInterface interface {
	Execute(command *getPurchasesCommand) (*[]dto.TicketOutputDTO, error)
}