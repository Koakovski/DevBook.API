package model

import (
	"errors"
	"time"
)

type Publication struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorId  uint64    `json:"authorId,omitempty"`
	Author    User      `json:"author,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (publication *Publication) Validate() error {
	if publication.Title == "" {
		return errors.New("title is missing")
	}

	if publication.Content == "" {
		return errors.New("content is missing")
	}

	return nil
}
