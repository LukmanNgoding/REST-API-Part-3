package delivery

import (
	"main.go/features/logistic/domain"
)

type RegisterFormat struct {
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Hp          string `json:"hp" form:"hp"`
	VehicleType string `json:"vehicle_type" form:"vehicle_type"`
}

type UpdateFormat struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Hp          string `json:"hp" form:"hp"`
	VehicleType string `json:"vehicle_type" form:"vehicle_type"`
}

type MyVendorFormat struct {
	Hp string `json:"hp" form:"hp"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Name: cnv.Name, Category: cnv.Category, Hp: cnv.Hp, VehicleType: cnv.VehicleType}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Name: cnv.Name, Category: cnv.Category, Hp: cnv.Hp, VehicleType: cnv.VehicleType}
	case MyVendorFormat:
		cnv := i.(MyVendorFormat)
		return domain.Core{Hp: cnv.Hp}
	}
	return domain.Core{}
}
