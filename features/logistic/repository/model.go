package repository

import (
	"main.go/features/logistic/domain"

	"gorm.io/gorm"
)

type Logistic struct {
	gorm.Model
	Name        string
	Category    string
	Hp          string
	VehicleType string
}

func FromDomain(du domain.Core) Logistic {
	return Logistic{
		Model:       gorm.Model{ID: du.ID},
		Name:        du.Name,
		Category:    du.Category,
		Hp:          du.Hp,
		VehicleType: du.VehicleType,
	}
}

func ToDomain(u Logistic) domain.Core {
	return domain.Core{
		ID:          u.ID,
		Name:        u.Name,
		Category:    u.Category,
		Hp:          u.Hp,
		VehicleType: u.VehicleType,
	}
}

func ToDomainArray(au []Logistic) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Name: val.Name, Category: val.Category, Hp: val.Hp, VehicleType: val.VehicleType})
	}
	return res
}
