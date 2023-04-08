package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	get_employees_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
)

type CompanyHandler struct {
	app *di.App
	logger logger.Logger
}

func NewCompanyHandler(
	logger logger.Logger,
	app *di.App,
) *CompanyHandler {
	return &CompanyHandler{
		logger: logger,
		app: app,
	}
}

// Customer godoc
// @Summary      Login as company
// @Description  Login as company with email and password
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        request   				body    dto.LoginInputDTO  true  "Login info"
// @Success      200  										{object}   	dto.CompanyAccountOutputDTO
// @Failure      404
// @Router       /company [post]
func (h *CompanyHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CompanyHandler.Login: Request received")
	var input dto.LoginInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Errorf("CompanyHandler.Login: Error at decoding request body, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := login_command.With(input.Email, input.Password)
	output, err := h.app.LoginAsCompanyUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CompanyHandler.Login: Error at searching for company account, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CompanyHandler.Login: New connection to account %s", input.Email)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

// Company godoc
// @Summary      Get all employees
// @Description  Get all employees based on company account id
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id   							path     		int true  "Company ID"
// @Success      200  										{object}   	[]dto.EmployeeOutputDTO
// @Failure      404
// @Router       /company/{id}/employees [get]
func (h *CompanyHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CompanyHandler.GetEmployees: Request received")
	
	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("CompanyHandler.GetEmployees: Unable to process request, %s", err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_employees_command.With(companyId)
	output, err := h.app.GetEmployeesUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CompanyHandler.GetEmployees: Unable to get employees, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CompanyHandler.GetEmployees: Employees infos delievered")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}