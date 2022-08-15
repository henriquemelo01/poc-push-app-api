package repository

import "poc-push-app-api/model"

type ReportsRepository interface {
	GetAll() ([]model.ReportModel, error)

	Create(report model.ReportModel) (model.ReportModel, error)

	GetById(id string) (model.ReportModel, error)

	Delete(id string) error
}
