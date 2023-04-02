package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	dto "github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
	"github.com/maicongiehl/techtur-api/internal/usecase"
)

type TravelHandler struct {
	Ctx context.Context
	TravelRepository *database.TravelRepository
}

func NewTravelHandler(ctx context.Context, travelRepository *database.TravelRepository) *TravelHandler {
	return &TravelHandler{
		Ctx: ctx,
		TravelRepository: travelRepository,
	}
}

// CreateTravel godoc
// @Summary      			Add travel
// @Description  			Create new travel
// @Tags         			Travel
// @Accept       			json
// @Produce      			json
// @Param        			request   				body     	dto.TravelInputDTO  			true  		"Travel Info"
// @Success      			201  												{object}   								object
// @Failure      			404
// @Router       			/travel [post]
func (h *TravelHandler) CreateTravel(w http.ResponseWriter, r *http.Request) {
	input, err := getTravelInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateTravelUseCase(*h.TravelRepository) 
	err = usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// TravelByCity godoc
// @Summary      			Get travel by user city
// @Description  			Use your customer credentials to enter in your account
// @Tags         			Travel
// @Accept       			json
// @Produce      			json
// @Param        			arrival_city_id   									path      		int  true  "Departure city"
// @Param        			departure_city_id   									path      		int  true  "Arrival city"
// @Success      			202  												{object}   		dto.TravelOutputDTO
// @Failure      			404
// @Router       			/travel/{departure_city_id}/{arrival_city_id}  [get]
func (h *TravelHandler) GetAllTraveslByDestiny(w http.ResponseWriter, r *http.Request) {
	dptCityId, err := strconv.Atoi(chi.URLParam(r, "departure_city_id"))
	if err != nil {
		returnErrMsg(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	arvCityId, err := strconv.Atoi(chi.URLParam(r, "arrival_city_id"))
	if err != nil {
		returnErrMsg(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := dto.TravelInputDTO{
		DepartureCityID: dptCityId,
		ArrivalCityID: arvCityId,
	}

	usecase := usecase.NewGetAllTravelsByDestinyUseCase(*h.TravelRepository)
	output, err := usecase.Execute(&input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	json.NewEncoder(w).Encode(output)
	
}


// DeleteTravel godoc
// @Summary      			Delete a specific travel
// @Description  			Get a travel
// @Tags         			Travel
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Travel Id"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/travel/{id} [delete]
func (h *TravelHandler) DeleteTravel(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if (err != nil || id <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	usecase := usecase.NewDeleteTravelUseCase(*h.TravelRepository)
	err = usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}


func getTravelInput(w http.ResponseWriter, r *http.Request) (*dto.TravelInputDTO, error) {
	var travel dto.TravelInputDTO
	err := json.NewDecoder(r.Body).Decode(&travel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &travel, err
	}
	return &travel, nil
}