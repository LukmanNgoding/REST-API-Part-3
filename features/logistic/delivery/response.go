package delivery

import "main.go/features/logistic/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type RegisterResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Hp          string `json:"hp"`
	VehicleType string `json:"vehicle_type"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Name: cnv.Name, Category: cnv.Category, Hp: cnv.Hp, VehicleType: cnv.VehicleType}
	case "all":
		var arr []RegisterResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, RegisterResponse{ID: val.ID, Name: val.Name, Category: val.Category, Hp: val.Hp, VehicleType: val.VehicleType})
		}
		res = arr
	case "del":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Name: cnv.Name, Category: cnv.Category, Hp: cnv.Hp, VehicleType: cnv.VehicleType}
	case "update":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Name: cnv.Name, Category: cnv.Category, Hp: cnv.Hp, VehicleType: cnv.VehicleType}
	}
	return res
}
