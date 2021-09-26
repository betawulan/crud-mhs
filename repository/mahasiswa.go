package repository

import (
	"context"
	"database/sql"
	"time"

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

func (m mahasiswaRepo) Fetch(ctx context.Context) ([]entity.Mahasiswa, error) {
	query := "SELECT id, name, email FROM mahasiswa"

	rows, err := m.db.QueryContext(ctx, query)
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
