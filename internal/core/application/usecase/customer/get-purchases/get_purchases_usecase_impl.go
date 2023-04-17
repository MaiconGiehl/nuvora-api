package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetPurchasesUsecase struct {
	ctx context.Context
	busPGSQLRepository *bus_entity.BusPGSQLRepository
	cityPGSQLRepository *city_entity.CityPGSQLRepository
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository
}

func NewGetPurchasesUsecase(
	ctx context.Context,
	busPGSQLRepository *bus_entity.BusPGSQLRepository,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *GetPurchasesUsecase {
	return &GetPurchasesUsecase{
		ctx: ctx,
		busPGSQLRepository: busPGSQLRepository,
		cityPGSQLRepository: cityPGSQLRepository,
		ticketPGSQLRepository: ticketPGSQLRepository,
		travelPGSQLRepository: travelPGSQLRepository,
	}	
}

func (u *GetPurchasesUsecase) Execute(
	command *getPurchasesCommand,
) (*[]dto.TicketOutputDTO, error) {
	var output []dto.TicketOutputDTO

	tickets, err := u.ticketPGSQLRepository.GetPurchases(command.accountId)
	if err != nil {
		return &output, err
	}

	travels, err := u.travelPGSQLRepository.FindAll()
	getTravels := func(id int) *travel_entity.Travel {
		for _, travel := range travels {
			if travel.ID == id {
				return travel
			}
		}
		return &travel_entity.Travel{}
	}

	allBus, err := u.busPGSQLRepository.FindAll()
	getBus := func (id int) *bus_entity.Bus {
		for _, bus := range allBus {
			if bus.ID == id {
				return bus
			}
		}
		return &bus_entity.Bus{}
	}


	allCities, _ := u.cityPGSQLRepository.FindAll()
	getCities := func(id int) city_entity.City {
		for _, city := range allCities {
			if city.ID == id {
				return *city
			}
		}
		return city_entity.City{}
	}


	for _, ticket := range *tickets {
		travel := getTravels(ticket.TravelID)
		bus := getBus(travel.BusID)
		departureCity := getCities(travel.Departure.CityID)
		arrivalCity := getCities(travel.Arrival.CityID)
		output = append(output, dto.TicketOutputDTO{
			ID: ticket.ID,
			StatusID: ticket.StatusID,
			TravelOutputDTO: *dto.NewTravelOutputDTO(
				travel.ID,
				travel.Price,
				bus.ID,
				bus.Number,
				bus.MaxPassengers,
				travel.AccountID,
				travel.Departure.Time,
				departureCity.Name,
				travel.Arrival.Time,
				arrivalCity.Name,
			),
			CreatedAt: ticket.CreatedAt,
		})
	}

	return &output, nil
}
