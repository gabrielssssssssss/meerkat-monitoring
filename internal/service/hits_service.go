package service

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/repository"
)

type HitService interface {
	Create(*models.Hit) error
	FindByToken(string) (*models.Hit, error)
	FindsByURL(string) ([]models.Hit, error)
}

type hitServiceImpl struct {
	repository repository.HitRepository
}

func NewHitService(repository repository.HitRepository) HitService {
	return &hitServiceImpl{
		repository: repository,
	}
}
