package service

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
)

func (s *transparencyServiceImpl) CreateDomainIndex() error {
	return s.repository.CreateDomainIndex()
}

func (s *transparencyServiceImpl) Create(transparency *models.Transparency) error {
	found, err := s.FindByDomain(transparency.Domain)
	if found != nil {
		return err
	}
	return s.repository.Create(transparency)
}

func (s *transparencyServiceImpl) FindByDomain(domain string) (*models.Transparency, error) {
	return s.repository.FindByDomain(domain)
}
