package handlers

import (
	"fmt"
	"net/http"

	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/shared/dto"
)

type CompanyHandler struct {
	app *di.App
}

func NewCompanyHandler(
	app *di.App,
) *CompanyHandler {
	return &CompanyHandler{
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
	var input dto.LoginInputDTO
	fmt.Print(input)
}
