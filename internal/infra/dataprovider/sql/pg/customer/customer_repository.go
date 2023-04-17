package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type CustomerPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
	logger logger.Logger
}

func NewCustomerPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *CustomerPGSQLRepository {
	return &CustomerPGSQLRepository{
		ctx: ctx,
		db: db,
		logger: logger,
	}
}

func (r *CustomerPGSQLRepository) FindCustomerByID(id int) (*Customer, error) {
	var output Customer

	stmt := `SELECT * FROM customer c WHERE c.id=$1`
	
	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.Cpf,
		&output.Name,
		&output.Phone,
		&output.Birth_date,
		&output.CompanyID,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("CustomerRepository.FindCustomerByID: Unable to find customer, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return &output, err
	}

	return &output, nil
}

func (r *CustomerPGSQLRepository) GetCustomersByCompanyID(id int) (*[]Customer, error) {
	var output []Customer

	stmt := "SELECT * FROM customer c WHERE c.company_id= $1"

	rows, err := r.db.Query(stmt, id)

	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Cpf,
			&customer.Name,
			&customer.Phone,
			&customer.Birth_date,
			&customer.CompanyID,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

		if err != nil {
			return &output, err
		}
		output = append(output, customer)
	}

	if err != nil {
		r.logger.Errorf("CustomerRepository.GetCustomersByCompanyID: Unable to find customers, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return &output, err
	}


	return &output, nil
}

func (r *CustomerPGSQLRepository) DeleteCustomerByID(customerId int, companyId int) error {
	stmt := `DELETE FROM customer c WHERE id = $1 AND c.company_id=$2`
	
	res, err := r.db.Exec(stmt, customerId, companyId)
	if err != nil {
		return err
	}

	affectedRowws, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRowws < 0 {
		return errors.New("none row affected")
	}

	return err
}

