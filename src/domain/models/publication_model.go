package model

import (
	"errors"
	"strings"
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

func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	if err := publication.format(); err != nil {
		return err
	}

	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("title is missing")
	}

	if publication.Content == "" {
		return errors.New("content is missing")
	}

	return nil
}

func (publication *Publication) format() error {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)

	return nil
}
