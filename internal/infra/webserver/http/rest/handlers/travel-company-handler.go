package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	create_travel_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/create-travel"
)

type TravelCompanyHandler struct {
	logger logger.Logger
	app    *di.App
}

func NewTravelCompanyHandler(
	logger logger.Logger,
	app *di.App,
) *TravelCompanyHandler {
	return &TravelCompanyHandler{
		logger: logger,
		app:    app,
	}
}

// Travel godoc
// @Summary      CreateTravel
// @Description  Create a new travel
// @Tags         TravelCompany
// @Accept       json
// @Produce      json
// @Param        id   				    path      int  true  "Travel company id"
// @Param        request   				body      dto.TravelInputDTO  true  "Travel info"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /travel-company/{id}/travels [post]
func (h *TravelCompanyHandler) CreateTravel(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("TravelCompanyHandler.CreateTravel: Request received")
	var input dto.TravelInputDTO

	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.CreateTravel: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.CreateTravel: Invalid body, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command, err := create_travel_command.With(companyId, input)
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.CreateTravel: Not what was expected, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = h.app.CreateTravelUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.CreateTravel: Error at creating travel, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("TravelCompanyHandler.Create: New travel created")
	w.WriteHeader(http.StatusCreated)
}
