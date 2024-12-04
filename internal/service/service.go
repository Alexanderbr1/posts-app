package service

import (
	"context"
	"posts-app/internal/config"
	"posts-app/internal/domain"
	"posts-app/internal/repository"
	"posts-app/pkg/cache"
	"time"
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	SignIn(ctx context.Context, input domain.SignInInput) (string, string, error)
	GenerateTokens(ctx context.Context, userID int) (string, string, error)
	ParseToken(token string) (int, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
}

type Post interface {
	Create(ctx context.Context, userId int, post domain.Post) (int, error)
	GetByID(ctx context.Context, id int) (domain.Post, error)
	GetAll(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, userId, id int, newPost domain.UpdatePost) error
	Delete(ctx context.Context, userId, id int) error
}

type Service struct {
	Authorization
	Post
}

func NewService(cfg *config.Config, cache *cache.MemoryCache, repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(cfg, repos.Authorization),
		Post:          NewPostService(cfg, cache, repos.Post),
	}
}
