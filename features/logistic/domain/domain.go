package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Name        string
	Category    string
	Hp          string
	VehicleType string
}

type Repository interface {
	InsertVendor(newVendor Core) (Core, error)
	UpdateVendor(newVendor Core) (Core, error)
	GetAllVendor() ([]Core, error)
	DeleteVendor(ID Core) error
}

type Service interface {
	AddVendor(newVendor Core) (Core, error)
	UpdateVendor(updatedData Core) (Core, error)
	ShowAllVendor() ([]Core, error)
	Delete(ID Core) error
	ExtractToken(c echo.Context) uint
}

type Handler interface {
	AddVendor() echo.HandlerFunc
	ShowAllVendor() echo.HandlerFunc
	UpdateVendor() echo.HandlerFunc
	DeleteVendor() echo.HandlerFunc
}
