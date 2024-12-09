package task6

import "go_testing/task6/dto"

type UserRepository interface {
	FindByID(id int) (dto.User, error)
	Create(dto.User) (int, error)
}

type UserService struct {
	r UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) FindByID(id int) (dto.User, error) {
	return s.r.FindByID(id)
}

func (s *UserService) Create(u dto.User) (int, error) {
	return s.r.Create(u)
}
