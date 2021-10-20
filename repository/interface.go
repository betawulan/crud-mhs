package repository

import (
	"context"

	"github.com/betawulan/crud-mhs/entity"
)

type MahasiswaRepository interface {
	Store(ctx context.Context, mahasiswa entity.Mahasiswa) (mhs entity.Mahasiswa, err error)
	Fetch(ctx context.Context, filter entity.FilterMahasiswa) ([]entity.Mahasiswa, error)
}
