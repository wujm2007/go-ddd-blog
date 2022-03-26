package model

import (
	"time"

	"go-ddd-blog/internal/service/post/domain/aggregate/post"
)

type CommentDTO struct {
	UserID int64 `json:"user_id"`

	ID      int64  `json:"id"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//goland:noinspection GoNameStartsWithPackageName
type PostDTO struct {
	UserID int64 `json:"user_id"`

	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`

	Comments []CommentDTO `json:"comments,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func BatchToDTO(entities []post.Post) []PostDTO {
	l := make([]PostDTO, len(entities))
	for i, e := range entities {
		l[i] = ToDTO(e)
	}
	return l
}

func ToDTO(entity post.Post) PostDTO {
	dto := PostDTO{
		UserID:    entity.UserID(),
		ID:        entity.ID(),
		Title:     entity.Title(),
		Content:   entity.Content(),
		Comments:  BatchCommentToDTO(entity.Comments()),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}

	return dto
}

func BatchCommentToDTO(entities []post.Comment) []CommentDTO {
	l := make([]CommentDTO, len(entities))
	for i, e := range entities {
		l[i] = CommentToDTO(e)
	}
	return l
}

func CommentToDTO(entity post.Comment) CommentDTO {
	return CommentDTO{
		UserID:    entity.UserID(),
		ID:        entity.ID(),
		Content:   entity.Content(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}
}

func CommentToEntity(dto CommentDTO) post.Comment {
	return *post.UnmarshalComment(dto.UserID, dto.ID, dto.Content, dto.CreatedAt, dto.UpdatedAt)
}
