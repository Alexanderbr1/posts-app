package service

import (
	"context"
	"errors"
	"fmt"
	"posts-app/internal/config"
	"posts-app/internal/domain"
	"posts-app/internal/repository"
	"posts-app/pkg/cache"
)

type PostService struct {
	cfg   *config.Config
	cache *cache.MemoryCache
	repo  repository.Post
}

func NewPostService(cfg *config.Config, cache *cache.MemoryCache, repo repository.Post) *PostService {
	return &PostService{cfg, cache, repo}
}

func (s *PostService) Create(ctx context.Context, userId int, post domain.Post) (int, error) {
	id, err := s.repo.Create(ctx, userId, post)
	if err != nil {
		return 0, err
	}

	s.cache.Set(fmt.Sprintf("%d.%d", id), domain.Post{ID: id, UserID: userId, Title: post.Title, Description: post.Description}, s.cfg.Cache.Ttl)

	return id, nil
}

func (s *PostService) GetByID(ctx context.Context, id int) (domain.Post, error) {
	post, err := s.cache.Get(fmt.Sprintf("%d.%d", id))
	if err == nil {
		return post.(domain.Post), nil
	}

	post, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}
	s.cache.Set(fmt.Sprintf("%d.%d", id), post, s.cfg.Cache.Ttl)

	return post.(domain.Post), nil
}

func (s *PostService) GetAll(ctx context.Context) ([]domain.Post, error) {
	return s.repo.GetAll(ctx)
}

func (s *PostService) Update(ctx context.Context, userId, id int, newPost domain.UpdatePost) error {
	if !newPost.IsValid() {
		return errors.New("update structure has no values")
	}

	s.cache.Set(fmt.Sprintf("%d.%d", userId, id), domain.Post{ID: id, Title: *newPost.Title, Description: *newPost.Description}, s.cfg.Cache.Ttl)

	return s.repo.Update(ctx, userId, id, newPost)
}

func (s *PostService) Delete(ctx context.Context, userId, id int) error {
	return s.repo.Delete(ctx, userId, id)
}
