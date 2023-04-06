package usecase

import (
	"context"

	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetPossibleTravelsUseCase struct {
	ctx context.Context
	travelPGSQLRepository *entity.TravelPGSQLRepository
}

func NewGetPossibleTravelsUseCase(
	ctx context.Context,
	travelPGSQLRepository *entity.TravelPGSQLRepository,
) *GetPossibleTravelsUseCase {
	return &GetPossibleTravelsUseCase{
		ctx: ctx,
		travelPGSQLRepository: travelPGSQLRepository,
	}	
}

func (u *GetPossibleTravelsUseCase) Execute(
	command *getPossibleTravelsCommand,
) (*[]entity.Travel, error) {
	var output []entity.Travel

	return &output, nil
}