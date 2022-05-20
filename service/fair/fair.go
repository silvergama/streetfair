package fair

import "github.com/silvergama/streetfair/entity"

type Repository interface {
	Create(f *entity.Fair) (int, error)
	Get(neighborhood string) ([]*entity.Fair, error)
	Update(f *entity.Fair) (int64, error)
	Delete(id int) error
}

type UseCase interface {
	CreateFair(fair *entity.Fair) (int, error)
	UpdateFair(fair *entity.Fair) (int, error)
	DeleteFair(id int) error
	GetFair(neighborhood string) ([]*entity.Fair, error)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateFair(fair *entity.Fair) (int, error) {
	return s.repo.Create(fair)
}

func (s *Service) GetFair(neighborhood string) ([]*entity.Fair, error) {
	return s.repo.Get(neighborhood)
}

func (s *Service) UpdateFair(fair *entity.Fair) (int, error) {
	rowsAffected, err := s.repo.Update(fair)
	return int(rowsAffected), err
}

func (s *Service) DeleteFair(id int) error {
	return s.repo.Delete(id)
}
