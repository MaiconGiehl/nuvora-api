package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/MaiconGiehl/API/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type CityHandler struct {
	Ctx context.Context
	CityRepository *database.CityRepository
}

func NewCityHandler(ctx context.Context, cityRepository *database.CityRepository) *CityHandler {
	return &CityHandler{
		Ctx: ctx,
		CityRepository: cityRepository,
	}
}

// CreateCity godoc
// @Summary      			Add city
// @Description  			Create new city
// @Tags         			City
// @Produce      			json
// @Param        			name   										path      	string  true  "City Name"
// @Success      			201  											{object}   	object
// @Failure      			404
// @Router       			/city/{name} [post]
func (h *CityHandler) CreateCity(w http.ResponseWriter, r *http.Request) {
	input := dto.CityInputDTO{
		Name: strings.ToUpper(chi.URLParam(r, "name")),
	}

	usecase := usecase.NewCreateCityUseCase(*h.CityRepository) 
	err := usecase.Execute(&input)
	if err != nil {
		returnErrMsg(w, err)
		returnErrMsg(w, err)
		return
	}

	json.NewEncoder(w).Encode("City created")
	w.WriteHeader(http.StatusCreated)
}


// DeleteCity godoc
// @Summary      			Delete a specific city
// @Description  			Delete a city
// @Tags         			City
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "City Id"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/city/{id} [delete]
func (h *CityHandler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	
	usecase := usecase.NewDeleteCityUseCase(*h.CityRepository)
	err = usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}