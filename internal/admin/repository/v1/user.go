package v1

import (
	"github.com/GodYao1995/Goooooo/internal/domain"
)

type UserRepository struct {
}

func NewUserRepository() domain.UserRepositoryFace {
	return &UserRepository{}
}
