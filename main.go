package main

import (
	"database/sql"
	"log"

	"github.com/betawulan/crud-mhs/delivery"
	"github.com/betawulan/crud-mhs/repository"
	"github.com/betawulan/crud-mhs/service"
	_ "github.com/go-sql-driver/mysql" // auto load pertama kali
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := sql.Open("mysql", "beta:12345beta@tcp(127.0.0.1:3306)/kuliah")
	if err != nil {
		log.Fatal("tdk dpt connect database")
	}

	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	delivery.RegisterMahasiswaRoute(mahasiswaService, e)

	e.Logger.Fatal(e.Start(":8000"))
}
