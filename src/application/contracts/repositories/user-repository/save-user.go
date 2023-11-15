package userrepository

type SaveUserRepository interface {
	Save(username string, email string, password string) (*CreateUserOutput, error)
}

type CreateUserOutput struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
