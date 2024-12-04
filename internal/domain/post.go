package domain

type Post struct {
	ID          int    `json:"-" db:"id"`
	UserID      int    `json:"-" db:"user_id"`
	Title       string `json:"title,max=55" db:"title" example:"Title"`
	Description string `json:"description,max=255" db:"description" example:"Description"`
}

type UpdatePost struct {
	Title       *string `json:"title" example:"Title"`
	Description *string `json:"description" example:"Description!"`
}

func (un UpdatePost) IsValid() bool {
	return un.Title != nil && un.Description != nil
}
