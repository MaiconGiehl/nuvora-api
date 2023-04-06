package rest

import (
	"net/http"

	handler_functions "github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	// /configs/di"
)

func LoginAsCustomer(w http.ResponseWriter, r *http.Request, customerHandler *handler_functions.CustomerHandler) {
	customerHandler.Login(w, r)
}