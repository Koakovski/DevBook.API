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
func (publicationRepository publicationRepository) FindById(publicationId uint64) (model.Publication, error) {
	var publication model.Publication

	rows, err := publicationRepository.db.Query(`
		SELECT p.id, p.title, p.content, p.likes, p.authorId, p.createdAt, 
		u.id, u.name, u.nickName, u.email, u.createdAt FROM publications p
		INNER JOIN users u
		ON u.id = p.authorId 
		WHERE p.id = ?
	`, publicationId,
	)
	if err != nil {
		return publication, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.AuthorId,
			&publication.CreatedAt,
			&publication.Author.ID,
			&publication.Author.Name,
			&publication.Author.NickName,
			&publication.Author.Email,
			&publication.Author.CreatedAt,
		); err != nil {
			return publication, err
		}
	}

	return publication, nil
}
