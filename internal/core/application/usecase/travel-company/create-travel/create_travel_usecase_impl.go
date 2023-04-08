package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type CreateTravelUsecase struct {
	ctx context.Context
	travelPGSQLRepository *entity.TravelPGSQLRepository
}

func NewCreateTravelUseCase(
	ctx context.Context,
	travelPGSQLRepository *entity.TravelPGSQLRepository,
) *CreateTravelUsecase {
	return &CreateTravelUsecase{
		ctx: ctx,
		travelPGSQLRepository: travelPGSQLRepository,
	}	
}

func (u *CreateTravelUsecase) Execute(
	command *createTravelCommand,
) (*[]dto.TravelOutputDTO, error) {
	var output []dto.TravelOutputDTO

	err := u.travelPGSQLRepository.CreateTravel()
	if err != nil {
		return &output, err
	}

	return &output, nil
}