package delivery

import (
	"net/http"
	"strconv"

	"github.com/betawulan/crud-mhs/entity"
	"github.com/betawulan/crud-mhs/service"

	"github.com/labstack/echo/v4"
)

type mahasiswaDelivery struct {
	mahasiswaService service.MahasiswaService
}

func (m mahasiswaDelivery) store(c echo.Context) error {
	var mahasiswa entity.Mahasiswa
	// merubah json request ke dalam struct
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
	filter := entity.FilterMahasiswa{}

	// query param itu optional
	// param & query param itu different
	filter.Name = c.QueryParam("name")
	filter.Email = c.QueryParam("email")
	filter.Order = c.QueryParam("order")

	limit := c.QueryParam("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		filter.Limit = uint64(limitInt)
	}

	page := c.QueryParam("page")
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		filter.Page = int(pageInt)
	}

	mahasiswas, err := m.mahasiswaService.Fetch(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, mahasiswas)

}

// func (m mahasiswaDelivery) getById(c echo.Context) error {
// 	id := c.Param("id")
// 	fmt.Println(id)
// 	return nil
// }

func RegisterMahasiswaRoute(mahasiswaService service.MahasiswaService, e *echo.Echo) {
	handler := mahasiswaDelivery{
		mahasiswaService: mahasiswaService,
	}

	e.POST("/mahasiswa", handler.store)
	e.GET("/mahasiswa", handler.fetch)
	// e.GET("/mahasiswa/:id/", handler.getById)
}
