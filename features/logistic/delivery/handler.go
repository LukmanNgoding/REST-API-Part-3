package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main.go/features/logistic/domain"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	o := e.Group("/vendors")
	o.Use(middleware.JWT([]byte("Anakmama!!12")))
	o.GET("", handler.ShowAllVendor())
	o.POST("", handler.AddVendor())
	o.POST("/update", handler.UpdateVendor())
	o.POST("/delete", handler.DeleteVendor())
}

func (bs *bookHandler) AddVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.AddVendor(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}
}

func (bs *bookHandler) UpdateVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.UpdateVendor(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "reg")))
	}
}

func (bs *bookHandler) ShowAllVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		res, err := bs.srv.ShowAllVendor()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success get book", ToResponse(res, "all")))
	}
}

func (bs *bookHandler) DeleteVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := bs.srv.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		err := bs.srv.Delete(cnv)
		if err != nil {
			return c.JSON(http.StatusOK, SuccessResponse("Berhasil Delete", err))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("Gagal Delete"))
	}
}
