package service

import "github.com/gabrielssssssssss/meerkat-monitoring/internal/models"

func (s *hitServiceImpl) Create(hit *models.Hit) error {
	return s.repository.Create(hit)
}

func (s *hitServiceImpl) FindByToken(token string) (*models.Hit, error) {
	return s.repository.FindByToken(token)
}

func (s *hitServiceImpl) FindsByURL(url string) ([]models.Hit, error) {
	return s.repository.FindsByURL(url)
}
