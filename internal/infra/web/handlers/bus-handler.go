package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type BusHandler struct {
	Ctx context.Context
	Db *sql.DB
}

func NewBusHandler(ctx context.Context, db *sql.DB) *BusHandler {
	return &BusHandler{
		Ctx: ctx,
		Db: db,
	}
}

// CreateBus godoc
// @Summary      			Add bus
// @Description  			Create new bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			request   				body      dto.BusInputDTO  true  "Bus Info"
// @Success      			200  											{object}   object
// @Failure      			404
// @Router       			/bus [post]
func (h *BusHandler) CreateBus(w http.ResponseWriter, r *http.Request) {
	busInputDTO, err := getBusInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	busRepository := database.NewBusRepository(h.Db)
	err = busRepository.Save(busInputDTO)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllBus godoc
// @Summary      			Get all bus
// @Description  			Get all bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Success      			200  						{object}   []dto.BusOutputDTO
// @Failure      			404
// @Router       			/bus [get]
func (h *BusHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	busRepository := database.NewBusRepository(h.Db)
	allBus, err := busRepository.GetAll()
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&allBus)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}

// GetBusById godoc
// @Summary      			Search for a specific bus
// @Description  			Get a bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      	int  true  "Bus ID"
// @Success      			200  										{object}   	dto.BusOutputDTO
// @Failure      			404
// @Router       			/bus/{id} [get]
func (h *BusHandler) GetBus(w http.ResponseWriter, r *http.Request) {
	id, err := getBusId(w, r)
	if err != nil {
		return
	}

	busRepository := database.NewBusRepository(h.Db)
	bus, err := busRepository.GetById(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&bus)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}

// DeleteBus godoc
// @Summary      			Delete a specific bus
// @Description  			Get a bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Bus Id"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/bus/{id} [delete]
func (h *BusHandler) DeleteBus(w http.ResponseWriter, r *http.Request) {
	id, err := getBusId(w, r)
	if err != nil {
		return
	}
	
	busRepository := database.NewBusRepository(h.Db)
	
	err = busRepository.Delete(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}


// UpdateBus godoc
// @Summary      			Delete a specific bus
// @Description  			Get a bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Account ID"
// @Param        			request   								body      		dto.BusInputDTO  true  "Bus info"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/bus/{id} [patch]
func (h *BusHandler) UpdateBus(w http.ResponseWriter, r *http.Request) {
	id, err := getBusId(w, r)
	if err != nil {
		return
	}

	bus, err := getBusInput(w, r)
	if err != nil {
		return
	}

	busRepository := database.NewBusRepository(h.Db)
	err = busRepository.Update(id, bus)
	
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}

func getBusId(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if (err != nil || id <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		return id, err
	}
	return id, nil
}

func getBusInput(w http.ResponseWriter, r *http.Request) (*dto.BusInputDTO, error) {
	var bus dto.BusInputDTO
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &bus, err
	}
	return &bus, nil
}

func returnErrMsg(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(msg)
}