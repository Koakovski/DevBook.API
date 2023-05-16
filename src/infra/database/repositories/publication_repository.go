package repository

import (
	"database/sql"
	model "devbook-api/src/domain/models"
)

type publicationRepository struct {
	db *sql.DB
}

func GetPublicationRepository(db *sql.DB) *publicationRepository {
	return &publicationRepository{db}
}

func (publicationRepository publicationRepository) Create(publicationModel model.Publication) (uint64, error) {
	statment, err := publicationRepository.db.Prepare(
		"INSERT INTO publications (title, content, authorId) VALUES (?,?,?)",
	)

	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(
		publicationModel.Title,
		publicationModel.Content,
		publicationModel.AuthorId,
	)
	if err != nil {
		return 0, err
	}

	publicationId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(publicationId), nil
}
