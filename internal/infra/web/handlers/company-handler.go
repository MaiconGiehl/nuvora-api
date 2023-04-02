package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	dto "github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
	"github.com/maicongiehl/nuvera-api/internal/usecase"
)

type CompanyHandler struct {
	Ctx context.Context
	CompanyRepository 			*database.CompanyRepository
	PersonRepository 				*database.PersonRepository
	AccountRepository 			*database.AccountRepository
	CustomerRepository      *database.CustomerRepository
}

func NewCompanyHandler(
	ctx context.Context,
	personRepository *database.PersonRepository,
	companyRepository *database.CompanyRepository,
	accountRepository *database.AccountRepository,
	customerRepository *database.CustomerRepository,
	) *CompanyHandler {
	return &CompanyHandler{
		Ctx: ctx,
		CompanyRepository: 			companyRepository,
		PersonRepository: 			personRepository,
		AccountRepository: 			accountRepository,
		CustomerRepository:     customerRepository,
	}
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

// Login godoc
// @Summary      			Login as company
// @Description  			Use your company credentials to enter in your account
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Param        			email   										path      		string  true  "Company email"
// @Param        			password   									path      		string  true  "Company password"
// @Success      			202  												{object}   		dto.CompanyAccountOutputDTO
// @Failure      			404
// @Router       			/company/{email}/{password} [get]
func (h *CompanyHandler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := h.getLoginInfo(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetCompanyAccountUseCase(*h.AccountRepository) 
	output, err := usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

// CreateEmployee godoc
// @Summary      			Add customer
// @Description  			Create new customer
// @Tags         			Company
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Company id"
// @Param        			request   				body      dto.CustomerAccountInputDTO  true  "Customer Info"
// @Success      			200  											{object}   object
// @Failure      			404
// @Router       			/company/{id}/employees [post]
func (h *CompanyHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	input, err := h.getCreateInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateCustomerAccountUseCase(*h.CustomerRepository, *h.PersonRepository, *h.AccountRepository) 
	err = usecase.Execute(input, companyId)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New customer created")
}

func (h *CompanyHandler) getLoginInfo(w http.ResponseWriter, r *http.Request) (*dto.CompanyLoginDTO, error) {
	return &dto.CompanyLoginDTO{
		Email: chi.URLParam(r, "email"),
		Password: chi.URLParam(r, "password"),
	}, nil
}

func (h *CompanyHandler) getCreateInput(w http.ResponseWriter, r *http.Request) (*dto.CustomerAccountInputDTO, error) {
	var customerAccount dto.CustomerAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&customerAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &customerAccount, err
	}
	return &customerAccount, nil
}
