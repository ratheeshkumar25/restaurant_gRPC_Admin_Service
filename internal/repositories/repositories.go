package repositories

import (
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/model"
	inter "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInter {
	return &AdminRepo{
		DB: db,
	}
}

func (a *AdminRepo) FindAdmin(user string) (*model.AdminModel, error) {
	var admin model.AdminModel
	err := a.DB.Where("username = ? ", user).First(&admin).Error
	return &admin, err
}
