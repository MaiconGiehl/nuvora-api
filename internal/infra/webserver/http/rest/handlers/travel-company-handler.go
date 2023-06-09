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
	delete_travel_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/delete-travel"
	get_all_bus_usecase_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/get-all-bus"
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
// @Security ApiKeyAuth
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

// Travel godoc
// @Summary      DeleteTravel
// @Description  DeleteTravel
// @Tags         TravelCompany
// @Accept       json
// @Produce      json
// @Param        id   				    path      int  true  "Travel  id"
// @Param        travelId   				    path      int  true  "Travel company id"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /travel-company/{id}/travel/{travelId} [delete]
// @Security ApiKeyAuth
func (h *TravelCompanyHandler) DeleteTravel(w http.ResponseWriter, r *http.Request) {
	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.DeleteTravel: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	travelId, err := strconv.Atoi(chi.URLParam(r, "travelId"))
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.DeleteTravel: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := delete_travel_command.With(travelId, companyId)

	err = h.app.DeleteTravelUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.DeleteTravel: Unable to get bus, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("TravelCompanyHandler.DeleteTravel: delete travel")
	w.WriteHeader(http.StatusAccepted)
}

// Travel godoc
// @Summary      GetAllBus
// @Description  GetAllBus
// @Tags         TravelCompany
// @Accept       json
// @Produce      json
// @Param        id   				    path      int  true  "Travel company id"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /travel-company/{id}/bus [get]
// @Security ApiKeyAuth
func (h *TravelCompanyHandler) GetAllBus(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("TravelCompanyHandler.GetAllBus: Request received")

	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.GetAllBus: Unable to process request, %s", err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_all_bus_usecase_command.With(companyId)
	output, err := h.app.GetAllBusUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("TravelCompanyHandler.GetAllBus: Unable to get bus, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("TravelCompanyHandler.GetAllBusTickets: bus infos delievered")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}
