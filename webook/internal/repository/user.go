package repository

import (
	"context"
	"go-study/webook/internal/domain"
	"go-study/webook/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

// UserRepository 存储
type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (u *UserRepository) FindUserByID(int64) {
	// 查询cache
	// 查询数据库
	// 回写cache
}

func (u *UserRepository) Create(ctx context.Context, user domain.User) error {
	return u.dao.Insert(ctx,
		dao.User{
			Email:    user.Email,
			Password: user.Password,
		})
}
