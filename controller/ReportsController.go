package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"poc-push-app-api/domain/services"
	"poc-push-app-api/dto"
)

type ReportController struct {
	repository services.ReportsRepository
}

func CreateReportController(reportRepository services.ReportsRepository) *ReportController {
	return &ReportController{
		repository: reportRepository,
	}
}

func (rc ReportController) SetupRoute(mux *chi.Mux) {

	const ReportsRoot = "/reports"

	mux.Route(ReportsRoot, func(router chi.Router) {
		router.Get("/", rc.GetAll)
		router.Get("/{id}", rc.GetById)
		router.Post("/", rc.Create)
		router.Delete("/{id}", rc.Delete)
	})
}

func (rc *ReportController) GetAll(writer http.ResponseWriter, req *http.Request) {

	reports, repositoryErr := rc.repository.GetAll()

	if repositoryErr != nil {
		http.Error(writer, repositoryErr.Error(), http.StatusInternalServerError)
		return
	}

	reportsJson, serializeJsonErr := json.Marshal(reports)
	if serializeJsonErr != nil {

		return
	}

	if len(reports) == 0 {
		writer.WriteHeader(http.StatusNoContent)
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	writer.Header().Set("Content-Type", "application/json")

	_, _ = writer.Write(reportsJson)
}

func (rc *ReportController) GetById(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")

	report, repositoryErr := rc.repository.GetById(id)

	if repositoryErr != nil {
		http.Error(writer, repositoryErr.Error(), http.StatusInternalServerError)
		return
	}

	reportJson, serializeJsonErr := json.Marshal(report)
	if serializeJsonErr != nil {
		http.Error(writer, repositoryErr.Error(), http.StatusNoContent)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, _ = writer.Write(reportJson)
}

func (rc ReportController) Create(writer http.ResponseWriter, req *http.Request) {

	defer func() {
		if reqBodyCloseError := req.Body.Close(); reqBodyCloseError != nil {
			writer.WriteHeader(http.StatusBadRequest)
			errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", reqBodyCloseError)
			_, _ = writer.Write([]byte(errorMessage))
		}
	}()

	createReportDto := dto.CreateReportDto{}
	if jsonParseError := json.NewDecoder(req.Body).Decode(&createReportDto); jsonParseError != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", jsonParseError.Error())
		http.Error(writer, errorMessage, http.StatusBadRequest)
		return
	}

	createdReportModel, repositoryErr := rc.repository.Create(createReportDto)

	if repositoryErr != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", repositoryErr.Error())
		http.Error(writer, errorMessage, http.StatusInternalServerError)
		return
	}

	createdReportModelJson, jsonSerializeError := json.Marshal(createdReportModel)

	if jsonSerializeError != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", jsonSerializeError.Error())
		http.Error(writer, errorMessage, http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write(createdReportModelJson)
}

func (rc ReportController) Delete(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")

	if repositoryErr := rc.repository.Delete(id); repositoryErr != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", repositoryErr.Error())
		http.Error(writer, errorMessage, http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
