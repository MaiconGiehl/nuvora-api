package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type CompanyPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
	logger logger.Logger
}

func NewCompanyPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *CompanyPGSQLRepository {
	return &CompanyPGSQLRepository{
		ctx: ctx,
		db: db,
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

func (r *CompanyPGSQLRepository) FindAllTravelCompanies() ([]*Company, error) {
	var output []*Company

	stmt := "SELECT * FROM company WHERE c.id IN (SELECT c.id FROM account a JOIN person p ON a.person_id=p.id JOIN company c ON p.company_id=c.id WHERE c.company_type_id=1 )"

	rows, err := r.db.Query(stmt)

	if err != nil {
		r.logger.Errorf("CompanyPGSQLRepository.FindAll: Unable to find companies, %s", err)
		return output, err
	}

	for rows.Next() {
		var company Company
		err := rows.Scan(
			&company.ID,
			&company.Cnpj,
			&company.SocialReason,
			&company.FantasyName,
			&company.Phone,
			&company.CompanyTypeID,
			&company.CreatedAt,
			&company.UpdatedAt,
		)

		if err != nil {
			return output, err
		}
		output = append(output, &company)
	}

	return output, nil
}