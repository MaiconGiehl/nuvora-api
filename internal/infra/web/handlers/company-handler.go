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

// GetAllCompany godoc
// @Summary      			Get all company
// @Description  			Get all company
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Success      			200  						{object}   []dto.CompanyOutputDTO
// @Failure      			404
// @Router       			/company [get]
func (h *CompanyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	usecase := usecase.NewGetAllCompanyUseCase(*h.CompanyRepository)
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

// GetAllEmployees godoc
// @Summary      			Get all employees
// @Description  			Get all employees
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "City Id"
// @Success      			200  						{object}   []dto.EmployeeOutputDTO
// @Failure      			404
// @Router       			/company/{id}/employees [get]
func (h *CompanyHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	companyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetEmployees(*h.CompanyRepository)
	output, err := usecase.Execute(companyID)
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

// GetAllEmployees godoc
// @Summary      			Get all employees
// @Description  			Get all employees
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "company Id"
// @Success      			200  						{object}   []dto.EmployeesTicketsOutputDTO
// @Failure      			404
// @Router       			/company/{id}/employees/tickets [get]
func (h *CompanyHandler) GetEmployeesTickets(w http.ResponseWriter, r *http.Request) {
	companyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetAllEmployeesTicketsUseCase(*h.CompanyRepository)
	output, err := usecase.Execute(companyID)
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


func getCompanyInput(w http.ResponseWriter, r *http.Request) (*dto.CompanyInputDTO, error) {
	var company dto.CompanyInputDTO
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &company, err
	}
	return &company, nil
}