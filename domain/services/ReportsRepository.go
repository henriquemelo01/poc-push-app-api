package services

import (
	"poc-push-app-api/domain/model"
	"poc-push-app-api/dto"
)

type ReportsRepository interface {
	GetAll() ([]model.ReportModel, error)

	Create(report dto.CreateReportDto) (model.ReportModel, error)

	GetById(id string) (model.ReportModel, error)

	Delete(id string) error
}
