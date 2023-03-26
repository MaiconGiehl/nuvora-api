package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/MaiconGiehl/API/internal/usecase"
	"github.com/go-chi/chi/v5"
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
// @Param        			request   				body      dto.TravelInputDTO  true  "Travel Info"
// @Success      			201  											{object}   object
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
	id, err := getTravelId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	
	usecase := usecase.NewDeleteTravelUseCase(*h.TravelRepository)
	err = usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}


func getTravelId(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if (err != nil || id <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		return id, err
	}
	return id, nil
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