package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"posts-app/internal/domain"
	"posts-app/internal/repository/psql"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	GetUserID(ctx context.Context, username, password string) (int, error)
	CreateToken(ctx context.Context, token domain.RefreshSession) error
	GetToken(ctx context.Context, token string) (domain.RefreshSession, error)
}

type Post interface {
	Create(ctx context.Context, userId int, post domain.Post) (int, error)
	GetByID(ctx context.Context, id int) (domain.Post, error)
	GetAll(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, userId, id int, newPost domain.UpdatePost) error
	Delete(ctx context.Context, userId, id int) error
}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psql.NewAuthRepository(db),
		Post:          psql.NewPostRepository(db),
	}
}
