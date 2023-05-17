package controller

import (
	presenter "devbook-api/src/app/presenters"
	model "devbook-api/src/domain/models"
	"devbook-api/src/infra/auth"
	"devbook-api/src/infra/database"
	repository "devbook-api/src/infra/database/repositories"
	"errors"
	"net/http"
)

func PublicationCreateController(w http.ResponseWriter, r *http.Request) {
	requestingUserId, err := auth.ExtractUserId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusUnauthorized, err)
		return
	}

	var publicationModel model.Publication
	statusCode, err := GetDataFromBody(r, &publicationModel, true)
	if err != nil {
		presenter.ErrorPresenter(w, statusCode, err)
		return
	}
	publicationModel.AuthorId = requestingUserId

	if err := publicationModel.Prepare(); err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	publicationRepository := repository.GetPublicationRepository(db)
	publicationModel.ID, err = publicationRepository.Create(publicationModel)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	presenter.ReponsePresenter(w, http.StatusCreated, publicationModel)
}

func PublicationFindOneController(w http.ResponseWriter, r *http.Request) {
	publicationId, err := GetId(r)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.GetDbConnection()
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	publicationRepository := repository.GetPublicationRepository(db)
	publication, err := publicationRepository.FindById(publicationId)
	if err != nil {
		presenter.ErrorPresenter(w, http.StatusInternalServerError, err)
		return
	}

	if publication.ID == 0 {
		presenter.ErrorPresenter(w, http.StatusBadRequest, errors.New("no publication found"))
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, publication)
}