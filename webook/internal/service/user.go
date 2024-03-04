package service

import (
	"context"
	"go-study/webook/internal/domain"
	"go-study/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (u *UserService) Signup(ctx context.Context, user domain.User) error {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	return u.repo.Create(ctx, user)
}

func (u *UserService) Login(ctx context.Context, user domain.User) error {
	u.repo.FindUserByID()
}
