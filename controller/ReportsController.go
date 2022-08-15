package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"poc-push-app-api/model"
	"poc-push-app-api/repository"
)

type ReportController struct {
	repository repository.ReportsRepository
}

func CreateReportController(reportRepository repository.ReportsRepository) *ReportController {
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

	reqBody := req.Body

	reqBodyData, readBodyErr := io.ReadAll(reqBody)

	if readBodyErr != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", readBodyErr.Error())
		http.Error(writer, errorMessage, http.StatusInternalServerError)
		return
	}

	reportModel := &model.ReportModel{}
	if jsonParseErr := json.Unmarshal(reqBodyData, reportModel); jsonParseErr != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", jsonParseErr.Error())
		http.Error(writer, errorMessage, http.StatusBadRequest)
		return
	}

	_, repositoryErr := rc.repository.Create(*reportModel)

	if repositoryErr != nil {
		errorMessage := fmt.Sprintf("Houve um erro por aqui: %s", repositoryErr.Error())
		http.Error(writer, errorMessage, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write(reqBodyData)
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
