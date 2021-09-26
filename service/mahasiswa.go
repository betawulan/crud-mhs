package service

import (
	"context"

	"github.com/betawulan/crud-mhs/entity"
	"github.com/betawulan/crud-mhs/repository"
)

type mahasiswaService struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func (m mahasiswaService) Store(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	// return m.mahasiswaRepo.Store(ctx, mahasiswa)
	mahasiswa, err := m.mahasiswaRepo.Store(ctx, mahasiswa)
	if err != nil {
		return entity.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func (m mahasiswaService) Fetch(ctx context.Context) ([]entity.Mahasiswa, error) {
	mahasiswas, err := m.mahasiswaRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return mahasiswas, nil
}

func NewMahasiswaService(mahasiswaRepo repository.MahasiswaRepository) MahasiswaService {
	return mahasiswaService{
		mahasiswaRepo: mahasiswaRepo,
	}
}
