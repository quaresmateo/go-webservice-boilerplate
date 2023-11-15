package services

import (
	userrepository "github.com/quaresmateo/go-webservice-boilerplate/src/application/contracts/repositories/user-repository"
	"github.com/quaresmateo/go-webservice-boilerplate/src/domain/usecases"
)

type CreateUserService struct {
	UserRepository userrepository.SaveUserRepository
}

func (service *CreateUserService) Execute(input *usecases.CreateUserInput) (*usecases.CreateUserOutput, error) {
	user, err := service.UserRepository.Save(input.Username, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return &usecases.CreateUserOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
