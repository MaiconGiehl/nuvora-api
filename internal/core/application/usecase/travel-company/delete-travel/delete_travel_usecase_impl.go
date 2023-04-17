package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type DeleteTravelUseCase struct {
	ctx context.Context
	logger logger.Logger
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository
}

func NewDeleteTravelUseCase(
	ctx context.Context,
	logger logger.Logger,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *DeleteTravelUseCase {
	return &DeleteTravelUseCase{
		ctx: ctx,
		logger: logger,
		travelPGSQLRepository: travelPGSQLRepository,
	}
}

func (u *DeleteTravelUseCase) Execute(command *deleteTravelCommand) error {
	err := u.travelPGSQLRepository.DeleteTravelByID(command.travelID, command.travelCompanyID)
	if err != nil {
		return err
	}

	return nil
}
