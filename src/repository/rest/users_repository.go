package rest

import (
	"github.com/davidhwang-ij/bookstore_oauth-api/src/domain/users"
	"github.com/davidhwang-ij/bookstore_oauth-api/src/utils/errors"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	return nil, nil
}
