package alert

import (
	"github.com/aldisypu/go-huberi/internal/pkg/entities"
	"gorm.io/gorm"
)

type Service interface {
	FindById(db *gorm.DB, alert *entities.Alert, id any) error
	Update(db *gorm.DB, alert *entities.Alert) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FindById(db *gorm.DB, alert *entities.Alert, id any) error {
	return s.repository.FindById(db, alert, id)
}

func (s *service) Update(db *gorm.DB, alert *entities.Alert) error {
	return s.repository.Update(db, alert)
}
