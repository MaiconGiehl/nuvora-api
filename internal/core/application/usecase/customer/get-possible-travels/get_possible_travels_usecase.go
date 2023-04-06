package usecase

import entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"

type GetPossibleTravelsUseCaseInterface interface {
	Execute(command *getPossibleTravelsCommand) (*[]entity.Travel, error)
}