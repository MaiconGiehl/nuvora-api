package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MaiconGiehl/API/internal/entity"
)

type CompanyRepository struct {
	Db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{Db: db}
}

func (r *CompanyRepository) Save(input *entity.Company) (error) {
	fmt.Print(time.Now().Format("2006-01-02T15:04:05"))
	fmt.Print(time.Now().Format("AAAAAAAAAAAAAA"))

	stmt := "INSERT INTO public.company (phone, cnpj, social_reason, fantasy_name, company_type_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"


	rows, err := r.Db.Exec(stmt, &input.Phone, &input.Cnpj, &input.SocialReason, &input.FantasyName, &input.CompanyTypeId,
		time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompanyRepository) Delete(input *entity.Company) error {
	stmt := "DELETE FROM public.company WHERE id= $1"

	rows, err := r.Db.Exec(stmt, &input.ID)
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}