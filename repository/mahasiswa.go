package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/betawulan/crud-mhs/entity"
)

type mahasiswaRepo struct {
	db *sql.DB
}

func (m mahasiswaRepo) Store(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	mahasiswa.CreatedAt = time.Now()
	query := "INSERT INTO mahasiswa (name, email, created_at) VALUES(?, ?, ?)"

	res, err := m.db.ExecContext(ctx, query, mahasiswa.Name, mahasiswa.Email, mahasiswa.CreatedAt)
	if err != nil {
		return entity.Mahasiswa{}, err
	}

	mahasiswa.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func (m mahasiswaRepo) Fetch(ctx context.Context, filter entity.FilterMahasiswa) ([]entity.Mahasiswa, error) {
	order := "created_at desc"
	if filter.Order == "asc" {
		order = "created_at asc"
	}

	qSelect := sq.Select("id", "name", "email").From("mahasiswa").OrderBy(order)
	if filter.Limit != 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Email != "" {
		qSelect = qSelect.Where(sq.Eq{"email": filter.Email})
	}

	if filter.Name != "" {
		qSelect = qSelect.Where(sq.Eq{"name": filter.Name})
	}

	if filter.Page != 0 {
		qSelect = qSelect.Offset(uint64(filter.Page))
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	mahasiswas := make([]entity.Mahasiswa, 0)
	for rows.Next() {
		var mahasiswa entity.Mahasiswa

		err = rows.Scan(&mahasiswa.ID, &mahasiswa.Name, &mahasiswa.Email)
		if err != nil {
			return nil, err
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas, nil
}

func NewMahasiswaRepository(db *sql.DB) MahasiswaRepository {
	return mahasiswaRepo{
		db: db,
	}
}

// layer delivery
// layer service
// layer repository (db)
// layer entity
