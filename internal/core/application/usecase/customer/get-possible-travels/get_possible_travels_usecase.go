package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type GetPossibleTravelsUseCaseInterface interface {
	Execute(command *getPossibleTravelsCommand) (*[]dto.TravelOutputDTO, error)
}
