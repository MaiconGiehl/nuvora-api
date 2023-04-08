package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type CreateTravelUseCaseInterface interface {
	Execute(command *createTravelCommand) (*[]dto.TravelOutputDTO, error)
}