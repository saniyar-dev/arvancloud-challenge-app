package repository

import (
	"arvancloud-challenge-app/internal/domain"

	"gorm.io/gorm"
)

type RequestRepository interface {
	SaveRequest(request *domain.Request) error
}

type requestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{db: db}
}

func (r *requestRepository) SaveRequest(request *domain.Request) error {
	return r.db.Create(request).Error
}
