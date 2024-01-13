package alert

import (
	"github.com/aldisypu/go-huberi/internal/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(db *gorm.DB, alert *entities.Alert, id any) error
	Update(db *gorm.DB, alert *entities.Alert) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindById(db *gorm.DB, alert *entities.Alert, id any) error {
	return db.Where("id = ?", id).Take(alert).Error
}

func (r *repository) Update(db *gorm.DB, alert *entities.Alert) error {
	return db.Save(alert).Error
}
