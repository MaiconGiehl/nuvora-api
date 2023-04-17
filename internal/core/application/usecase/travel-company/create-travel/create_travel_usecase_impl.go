package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type CreateTravelUseCase struct {
	ctx                    context.Context
	logger                 logger.Logger
	busPGSQLRepository     *bus_entity.BusPGSQLRepository
	cityPGSQLRepository    *city_entity.CityPGSQLRepository
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	travelPGSQLRepository  *travel_entity.TravelPGSQLRepository
}

func NewCreateTravelUseCase(
	ctx context.Context,
	logger logger.Logger,
	busPGSQLRepository *bus_entity.BusPGSQLRepository,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *CreateTravelUseCase {
	return &CreateTravelUseCase{
		ctx:                    ctx,
		logger:                 logger,
		travelPGSQLRepository:  travelPGSQLRepository,
		busPGSQLRepository:     busPGSQLRepository,
		cityPGSQLRepository:    cityPGSQLRepository,
		companyPGSQLRepository: companyPGSQLRepository,
	}
}

func (u *CreateTravelUseCase) Execute(
	command *createTravelCommand,
) error {

	err := u.validateInput(command)
	if err != nil {
		customErr := fmt.Sprintf("invalid field: %s", err.Error())
		err = errors.New(customErr)
		u.logger.Errorf("CreateTravelUseCase.Execute: Unable to create travel, %s", err.Error())
		return err
	}

	err = u.travelPGSQLRepository.CreateTravel(
		command.CompanyID,
		command.Price,
		command.BusID,
		command.DepartureTime,
		command.DepartureCityID,
		command.ArrivalTime,
		command.ArrivalCityID,
	)

	return err
}

func (u *CreateTravelUseCase) validateInput(input *createTravelCommand) error {
	if input.Price <= 0 {
		return errors.New("price can't be lower or equal to zero")
	}

	now := time.Now()
	if input.ArrivalTime.Compare(input.ArrivalTime) == -1 || input.DepartureTime.Compare(input.ArrivalTime) == 0 {
		return errors.New("arrival time must be after departure")
	}

	if input.DepartureCityID == input.ArrivalCityID {
		return errors.New("arrival and departure city must be different")
	}

	if input.DepartureTime.Compare(now) == 1 || input.DepartureTime.Compare(now) == 0 {
		return errors.New("departure time must be now or in the future")
	}

	if input.ArrivalTime.Compare(input.DepartureTime) == -1 || input.ArrivalTime.Compare(input.DepartureTime) == 0 {
		return errors.New("arrival must be after departure")
	}

	_, err := u.companyPGSQLRepository.FindCompanyByID(input.CompanyID)
	if err != nil {
		return errors.New("company not found")
	}

	bus, err := u.busPGSQLRepository.FindBusByID(input.BusID)
	if err != nil || bus.CompanyID != input.CompanyID {
		return errors.New("bus not found")
	}

	_, err = u.cityPGSQLRepository.FindCityByID(input.DepartureCityID)
	if err != nil {
		return errors.New("city not found")
	}

	_, err = u.cityPGSQLRepository.FindCityByID(input.ArrivalCityID)
	if err != nil {
		return errors.New("city not found")
	}

	return nil
}
