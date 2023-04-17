package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	delete_employee_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/delete-employee"
	get_employees_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees"
	get_employees_tickets_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees-tickets"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
	pay_tickets_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/pay-tickets"
)

type CompanyHandler struct {
	app    *di.App
	logger logger.Logger
}

func NewCompanyHandler(
	logger logger.Logger,
	app *di.App,
) *CompanyHandler {
	return &CompanyHandler{
		logger: logger,
		app:    app,
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

	output.SetAccessToken(h.createJWT(r, output.PermissionLevel)) 

	h.logger.Infof("CompanyHandler.Login: New connection to account %s", input.Email)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

func (h *CompanyHandler) createJWT(r *http.Request, permission_level int) string {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"exp":              time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
		"permission_level": permission_level,
	})

	return tokenString
}

// Company godoc
// @Summary      Get all employees
// @Description  Get all employees based on company account id
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id   							path     		int true  "Company ID"
// @Success      202  										{object}   	[]dto.EmployeeOutputDTO
// @Failure      404
// @Router       /company/{id}/employees [get]
// @Security ApiKeyAuth
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

// Company godoc
// @Summary      GetEmployeesTickets
// @Description  GetEmployeesTickets
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id   							path     		int true  "Company ID"
// @Success      200  										{object}   	[]dto.EmployeeOutputDTO
// @Failure      404
// @Router       /company/{id}/employees/tickets [get]
// @Security ApiKeyAuth
func (h *CompanyHandler) GetEmployeesTickets(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CompanyHandler.GetEmployeesTickets: Request received")

	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("CompanyHandler.GetEmployeesTickets: Unable to process request, %s", err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_employees_tickets_command.With(companyId)
	output, err := h.app.GetEmployeesTicketsUsecase.Execute(command)
	if err != nil {
		h.logger.Errorf("CompanyHandler.GetEmployeesTickets: Unable to get employees, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CompanyHandler.GetEmployeesTickets: Employees infos delievered")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

// Company godoc
// @Summary      DeleteEmployee
// @Description  DeleteEmployee
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id   							path     		int true  "Company ID"
// @Param        employeeId   							path     		int true  "Employee ID"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /company/{id}/employee/{employeId} [delete]
// @Security ApiKeyAuth
func (h *CompanyHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CompanyHandler.DeleteEmployee: Request received")

	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("CompanyHandler.DeleteEmployee: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	employeeId, err := strconv.Atoi(chi.URLParam(r, "employeeId"))
	if err != nil {
		h.logger.Errorf("CompanyHandler.DeleteEmployee: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := delete_employee_command.With(employeeId, companyId)

	err = h.app.DeleteEmployeeUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CompanyHandler.DeleteEmployee: Unable to delete employee, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CompanyHandler.DeleteEmployee: employee deleted")
	w.WriteHeader(http.StatusAccepted)
}

// Company godoc
// @Summary      PayAllTickets
// @Description  PayAllTickets
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id   							path     		int true  "Company ID"
// @Success      200  										{object}   	[]dto.EmployeeOutputDTO
// @Failure      404
// @Router       /company/{id}/employees/tickets [patch]
// @Security ApiKeyAuth
func (h *CompanyHandler) PayAllTickets(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CompanyHandler.PayAllTickets: Request received")

	companyId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("CompanyHandler.PayAllTickets: Unable to process request, %s", err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := pay_tickets_command.With(companyId)
	rowsAffected, err := h.app.PayTickets.Execute(command)
	if err != nil {
		h.logger.Errorf("CompanyHandler.PayAllTickets: Unable to pay tickets, %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CompanyHandler.PayAllTickets: tickets paid")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(rowsAffected)
}
