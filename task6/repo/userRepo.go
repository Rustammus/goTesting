package repo

import "go_testing/task6/dto"

//go:generate mockgen -source=userRepo.go -destination=mock/mock.go

type UserRepository interface {
	FindByID(id int) (dto.User, error)
	Create(dto.User) (int, error)
}
