package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

// GetQuery godoc
// @Summary      Add bus
// @Description  Create new bus
// @Tags         Bus
// @Accept       json
// @Produce      json
// @Param        request   				body      dto.BusInputDTO  true  "BusInfo"
// @Success      200  											{object}   object
// @Failure      404
// @Router       /bus [post]
func (h *BusHandler) CreateBus(w http.ResponseWriter, r *http.Request) {
	var bus dto.BusInputDTO
	err := json.NewDecoder(r.Body).Decode(&bus)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	busRepository := database.NewBusRepository(h.Db)
	busRepository.Save(bus.Number, bus.MaxPassengers)

	w.WriteHeader(http.StatusCreated)
}

// GetQuery godoc
// @Summary      Get all bus
// @Description  Get all bus
// @Tags         Bus
// @Accept       json
// @Produce      json
// @Success      200  						{object}   []dto.BusOutputDTO
// @Failure      404
// @Router       /bus [get]
func (h *BusHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	busRepository := database.NewBusRepository(h.Db)
	allBus, err := busRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = json.NewEncoder(w).Encode(&allBus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
}

// GetQuery godoc
// @Summary      Search for a specific bus
// @Description  Get a bus
// @Tags         Bus
// @Accept       json
// @Produce      json
// @Param        id   			path      		int  true  "Account ID"
// @Success      200  										{object}   dto.BusOutputDTO
// @Failure      404
// @Router       /bus/{id} [get]
func (h *BusHandler) GetBus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	busId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	busRepository := database.NewBusRepository(h.Db)
	bus, err := busRepository.GetById(busId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = json.NewEncoder(w).Encode(&bus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
}

// GetQuery godoc
// @Summary      Delete a specific bus
// @Description  Get a bus
// @Tags         Bus
// @Accept       json
// @Produce      json
// @Param        id   			path      		int  true  "Account ID"
// @Success      200  										{object}   object
// @Failure      404
// @Router       /bus/{id} [delete]
func (h *BusHandler) DeleteBus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	busId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Print(err)
	}

	busRepository := database.NewBusRepository(h.Db)
	err = busRepository.Delete(busId)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
}


// GetQuery godoc
// @Summary      Delete a specific bus
// @Description  Get a bus
// @Tags         Bus
// @Accept       json
// @Produce      json
// @Param        id   			path      		int  true  "Account ID"
// @Param        request   				body      dto.BusInputDTO  true  "BusInfo"
// @Success      200  										{object}   object
// @Failure      404
// @Router       /bus/{id} [patch]
func (h *BusHandler) UpdateBus(w http.ResponseWriter, r *http.Request) {
	fmt.Print("a")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if (err != nil || id <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bus dto.BusInputDTO
	err = json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Print("a")

	busRepository := database.NewBusRepository(h.Db)
	err = busRepository.Update(id, &bus)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(err)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
}