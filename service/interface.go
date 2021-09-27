package service

import (
	"context"

	"github.com/betawulan/crud-mhs/entity"
)

type MahasiswaService interface {
	Store(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	Fetch(ctx context.Context) ([]entity.Mahasiswa, error)
}
