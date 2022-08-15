package di

import (
	"poc-push-app-api/controller"
	"poc-push-app-api/repository"
)

func CreateReportsRepository() *repository.FakeReportsRepositoryImpl {
	return repository.CreateFakeReportsRepositoryImpl()
}

func CreateReportsController() *controller.ReportController {
	reportsRepository := CreateReportsRepository()
	return controller.CreateReportController(reportsRepository)
}
