package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	dto "github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/MaiconGiehl/API/internal/usecase"
)

type CompanyHandler struct {
	Ctx context.Context
	CompanyRepository *database.CompanyRepository
}

func NewCompanyHandler(ctx context.Context, companyRepository *database.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{
		Ctx: ctx,
		CompanyRepository: companyRepository,
	}
}

// CreateCompany godoc
// @Summary      			Add company
// @Description  			Create new company
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Param        			request   				body      dto.CompanyInputDTO  true  "Company Info"
// @Success      			201  											{object}   object
// @Failure      			404
// @Router       			/company [post]
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	input, err := getCompanyInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateCompanyUseCase(*h.CompanyRepository) 
	err = usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getCompanyInput(w http.ResponseWriter, r *http.Request) (*dto.CompanyInputDTO, error) {
	var company dto.CompanyInputDTO
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &company, err
	}
	return &company, nil
}