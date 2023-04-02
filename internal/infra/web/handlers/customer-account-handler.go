package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	dto "github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
	"github.com/maicongiehl/nuvera-api/internal/usecase"
)

type CustomerAccountHandler struct {
	Ctx 										context.Context
	CustomerRepository 			*database.CustomerRepository
	PersonRepository 				*database.PersonRepository
	AccountRepository 			*database.AccountRepository
}

func NewCustomerAccountHandler(
	ctx 										context.Context,
	customerRepository 			*database.CustomerRepository,
	personRepository 				*database.PersonRepository,
	accountRepository 			*database.AccountRepository,
	) *CustomerAccountHandler {
	return &CustomerAccountHandler{
		Ctx: 									ctx,
		CustomerRepository: 	customerRepository,
		PersonRepository: 		personRepository,
		AccountRepository: 		accountRepository,
	}
}

// Login godoc
// @Summary      			Login as customer
// @Description  			Use your customer credentials to enter in your account
// @Tags         			Customer
// @Accept       			json
// @Produce      			json
// @Param        			email   										path      		string  true  "Customer email"
// @Param        			password   									path      		string  true  "Customer password"
// @Success      			202  												{object}   		dto.CustomerAccountOutputDTO
// @Failure      			404
// @Router       			/customer/{email}/{password} [get]
func (h *CustomerAccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := h.getLoginInfo(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetCustomerAccountUseCase(*h.AccountRepository) 
	output, err := usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}


func (h *CustomerAccountHandler) getCreateInput(w http.ResponseWriter, r *http.Request) (*dto.CustomerAccountInputDTO, error) {
	var customerAccount dto.CustomerAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&customerAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &customerAccount, err
	}
	return &customerAccount, nil
}

func (h *CustomerAccountHandler) getLoginInfo(w http.ResponseWriter, r *http.Request) (*dto.LoginCustomerInputDTO, error) {
	return &dto.LoginCustomerInputDTO{
		Email: chi.URLParam(r, "email"),
		Password: chi.URLParam(r, "password"),
	}, nil
}