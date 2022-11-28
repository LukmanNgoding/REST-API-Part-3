package repository

import (
	"errors"

	"main.go/features/logistic/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) InsertVendor(newVendor domain.Core) (domain.Core, error) {
	var cnv Logistic
	cnv = FromDomain(newVendor)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	newVendor = ToDomain(cnv)
	return newVendor, nil
}

func (rq *repoQuery) UpdateVendor(updatedData domain.Core) (domain.Core, error) {
	var cnv Logistic
	cnv = FromDomain(updatedData)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	updatedData = ToDomain(cnv)
	return updatedData, nil
}

func (rq *repoQuery) GetAllVendor() ([]domain.Core, error) {
	var resQry []Logistic
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) DeleteVendor(ID domain.Core) error {
	var res Logistic = FromDomain(ID)
	if err := rq.db.Where("id = ?", res.ID).Delete(&res).Error; err != nil {
		return errors.New("gagal delete")
	}
	// ID = ToDomain(res)
	return errors.New("berhasil delete")
}
