package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/maicongiehl/nuvora-api/configs/env"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	postgresdb "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
	logrus_config "github.com/maicongiehl/nuvora-api/internal/infra/log/logrus"

	"github.com/stretchr/testify/suite"
)

var ctx = context.Background()

type CreateTravelUseCaseImplTestSuite struct {
	suite.Suite
	ctx                    context.Context
	logger                 *logrus_config.LogrusLogger
	createTravelRepository *CreateTravelUseCase
}

func (suite *CreateTravelUseCaseImplTestSuite) SetupTest() {
	logger := logrus_config.NewLogrusLogger()

	suite.ctx = ctx
	suite.logger = logger
	env := env.LoadConfig("../../../../../../.env", logger)

	db := postgresdb.ConnectWithDB(
		logger,
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)

	companyPGSQLRepository := company_entity.NewCompanyPGSQLRepository(ctx, db, logger)
	busPGSQLRepository := bus_entity.NewBusPGSQLRepository(ctx, db, logger)
	cityPGSQLRepository := city_entity.NewCityPGSQLRepository(ctx, db, logger)
	travelPGSQLRepository := travel_entity.NewTravelPGSQLRepository(ctx, db, logger)

	useCaseRepository := NewCreateTravelUseCase(
		ctx,
		logger,
		busPGSQLRepository,
		cityPGSQLRepository,
		companyPGSQLRepository,
		travelPGSQLRepository,
	)

	suite.createTravelRepository = useCaseRepository

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CreateTravelUseCaseImplTestSuite))
}

func (s *CreateTravelUseCaseImplTestSuite) TestCreateTravel() {
	command, err := With(
		1,
		dto.TravelInputDTO{
			Price: 1,
			BusID: 1,
			Departure: dto.DepartureInputDTO{
				Time:   time.Now().Format("2006-01-02T00:00:00"),
				CityID: 1,
			},
			Arrival: dto.ArrivalInputDTO{
				Time:   time.Now().AddDate(10, 10, 10).Format("2006-01-02T00:00:00"),
				CityID: 2,
			},
		},
	)
	s.NoError(err)

	err = s.createTravelRepository.Execute(command)
	s.NoError(err)

	command, err = With(
		1000,
		dto.TravelInputDTO{
			Price: 0,
			BusID: -1,
			Departure: dto.DepartureInputDTO{
				Time:   time.Now().Format("2006-01-02T00:00:00"),
				CityID: -1,
			},
			Arrival: dto.ArrivalInputDTO{
				Time:   time.Now().Format("2006-01-02T00:00:00"),
				CityID: -1,
			},
		},
	)
	s.NoError(err)

	err = s.createTravelRepository.Execute(command)
	s.Error(err)

	_, err = With(
		1000,
		dto.TravelInputDTO{
			Departure: dto.DepartureInputDTO{
				Time: "ASDASDAS",
			},
			Arrival: dto.ArrivalInputDTO{
				Time: time.Now().Format("2006-01-02T00:00:00"),
			},
		},
	)
	s.Error(err)

	_, err = With(
		1000,
		dto.TravelInputDTO{
			Price: 0,
			BusID: -1,
			Departure: dto.DepartureInputDTO{
				Time:   time.Now().Format("2006-01-02T00:00:00"),
				CityID: 0,
			},
			Arrival: dto.ArrivalInputDTO{
				Time:   "ASDSDASD",
				CityID: 1,
			},
		},
	)
	s.Error(err)

}
