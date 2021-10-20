package entity

import "time"

type Mahasiswa struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
}

type FilterMahasiswa struct {
	Limit uint64 // sama seperti int64 tp bedanya uint64 positif saja
	Page  int
	Email string
	Name  string
	Order string
}
