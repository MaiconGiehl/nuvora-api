package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
)

type GetAllBusUseCase struct {
	ctx                context.Context
	logger             logger.Logger
	busPGSQLRepository *bus_entity.BusPGSQLRepository
}

func NewGetAllBusUseCase(
	ctx context.Context,
	logger logger.Logger,
	busPGSQLRepository *bus_entity.BusPGSQLRepository,
) *GetAllBusUseCase {
	return &GetAllBusUseCase{
		ctx:                ctx,
		logger:             logger,
		busPGSQLRepository: busPGSQLRepository,
	}
}

func (u *GetAllBusUseCase) Execute(command *getAllBusCommand) ([]*dto.BusOutputDTO, error) {
	var output []*dto.BusOutputDTO

	bus, err := u.busPGSQLRepository.FindBusByCompanyAccountID(command.travelCompanyID)
	if err != nil {
		return output, err
	}

	for _, b := range bus {
		output = append(output, dto.NewBusOutputDTO(
			b.ID,
			b.Number,
			b.MaxPassengers,
			b.CompanyID,
		))

	}

	return output, nil
}
