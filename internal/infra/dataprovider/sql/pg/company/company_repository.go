package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type CompanyPGSQLRepository struct {
	ctx    context.Context
	db     *sql.DB
	logger logger.Logger
}

func NewCompanyPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *CompanyPGSQLRepository {
	return &CompanyPGSQLRepository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *CompanyPGSQLRepository) FindCompanyByID(companyId int) (*Company, error) {
	var output Company
	stmt := `SELECT * FROM company c WHERE c.id =$1`

	row := r.db.QueryRow(stmt, companyId)

	err := row.Scan(
		&output.ID,
		&output.Cnpj,
		&output.SocialReason,
		&output.FantasyName,
		&output.Phone,
		&output.CompanyTypeID,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("CompanyRepository.FindCompanyByID: Unable to find company, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return &output, err
	}

	return &output, nil
}
