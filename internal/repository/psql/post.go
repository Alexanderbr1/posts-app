package psql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"posts-app/internal/domain"
	"strings"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, userId int, post domain.Post) (int, error) {
	var id int

	createPostQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id",
		postsTable)
	row := r.db.QueryRowContext(ctx, createPostQuery, userId, post.Title, post.Description)
	err := row.Scan(&id)

	return id, err
}

func (r *PostRepository) GetByID(ctx context.Context, id int) (domain.Post, error) {
	var post domain.Post

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	err := r.db.GetContext(ctx, &post, query, id)

	return post, err
}

func (r *PostRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post

	query := fmt.Sprintf("SELECT * FROM %s", postsTable)
	err := r.db.SelectContext(ctx, &posts, query)

	return posts, err
}

func (r *PostRepository) Update(ctx context.Context, userId, id int, newPost domain.UpdatePost) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if newPost.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *newPost.Title)
		argId++
	}

	if newPost.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *newPost.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d", postsTable, setQuery, argId, argId+1)
	args = append(args, id, userId)

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}

func (r *PostRepository) Delete(ctx context.Context, userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", postsTable)

	_, err := r.db.ExecContext(ctx, query, userId, id)

	return err
}
