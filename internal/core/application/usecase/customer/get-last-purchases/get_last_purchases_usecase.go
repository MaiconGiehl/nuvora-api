package usecase

import entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"

type GetLastPurchasesInterface interface {
	Execute(command *getLastPurchasesCommand) (*[]entity.Ticket, error)
}