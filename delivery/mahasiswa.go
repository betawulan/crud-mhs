package delivery

import (
	"net/http"

	"github.com/betawulan/crud-mhs/entity"
	"github.com/betawulan/crud-mhs/service"

	"github.com/labstack/echo/v4"
)

type mahasiswaDelivery struct {
	mahasiswaService service.MahasiswaService
}

func (m mahasiswaDelivery) store(c echo.Context) error {
	var mahasiswa entity.Mahasiswa
	err := c.Bind(&mahasiswa)
	if err != nil {
		err = c.JSON(http.StatusBadRequest, err)

		return err
	}

	mahasiswa, err = m.mahasiswaService.Store(c.Request().Context(), mahasiswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, mahasiswa)
}

func (m mahasiswaDelivery) fetch(c echo.Context) error {
	ctx := c.Request().Context()
	mahasiswas, err := m.mahasiswaService.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, mahasiswas)

}

func RegisterMahasiswaRoute(mahasiswaService service.MahasiswaService, e *echo.Echo) {
	handler := mahasiswaDelivery{
		mahasiswaService: mahasiswaService,
	}

	e.POST("/mahasiswa", handler.store)
	e.GET("/mahasiswa", handler.fetch)
}
