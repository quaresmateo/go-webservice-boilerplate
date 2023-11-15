package usecases

type CreateUserUseCase interface {
	Execute(input CreateUserInput) (CreateUserOutput, error)
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
