package main

import (
	_ "github.com/maicongiehl/nuvora-api/docs"

	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest"
)

//	@title			Nuvora API
//	@version		1.0
//	@description	Product API with auhtentication
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Nuvora
//	@contact.url	https://techtur.com.br
//	@contact.email	atendimento@nuvora.com.br

//	@license.name	Nuvora Promotora License
//	@license.url	https://nuvora.com.br

// @host		localhost:8080
// @BasePath  /nuvora/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	rest.StartServer()
}
