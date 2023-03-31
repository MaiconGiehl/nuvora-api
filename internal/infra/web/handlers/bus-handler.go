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

type BusHandler struct {
	Ctx context.Context
	BusRepository *database.BusRepository
}

func NewBusHandler(ctx context.Context, busRepository *database.BusRepository) *BusHandler {
	return &BusHandler{
		Ctx: ctx,
		BusRepository: busRepository,
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
	input, err := getBusInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateBusUseCase(*h.BusRepository) 
	err = usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New bus created")
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
	usecase := usecase.NewGetAllBusUseCase(*h.BusRepository)
	output, err := usecase.Execute()
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&output)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusFound)
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
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetBusUseCase(*h.BusRepository)
	output, err := usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&output)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusFound)

}

// DeleteBus godoc
// @Summary      			Delete a specific bus
// @Description  			Get a bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Bus Id"
// @Success      			202  										{object}   		object
// @Failure      			404
// @Router       			/bus/{id} [delete]
func (h *BusHandler) DeleteBus(w http.ResponseWriter, r *http.Request) {
	id, err := getBusId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	
	usecase := usecase.NewDeleteBusUseCase(*h.BusRepository)
	err = usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	
	w.WriteHeader(http.StatusAccepted)
}


// UpdateBus godoc
// @Summary      			Delete a specific bus
// @Description  			Get a bus
// @Tags         			Bus
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Account ID"
// @Param        			request   								body      		dto.BusInputDTO  true  "Bus info"
// @Success      			202  										{object}   		object
// @Failure      			404
// @Router       			/bus/{id} [patch]
func (h *BusHandler) UpdateBus(w http.ResponseWriter, r *http.Request) {
	id, err := getBusId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	bus, err := getBusInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewUpdateBusUseCase(*h.BusRepository)
	err = usecase.Execute(id, bus)
	
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
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
	json.NewEncoder(w).Encode(err.Error())
}