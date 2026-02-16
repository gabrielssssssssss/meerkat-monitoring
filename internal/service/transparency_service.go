package service

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/repository"
)

type TransparencyService interface {
	CreateDomainIndex() error
	Create(*models.Transparency) error
	FindByDomain(string) (*models.Transparency, error)
}

type transparencyServiceImpl struct {
	repository repository.TransparencyRepository
}

func NewTransparencyService(repository repository.TransparencyRepository) TransparencyService {
	return &transparencyServiceImpl{
		repository: repository,
	}
}
