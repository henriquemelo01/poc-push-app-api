package di

import (
	"awesomeProject/controller"
	"awesomeProject/repository"
)

func CreateReportsRepository() *repository.FakeReportsRepositoryImpl {
	return repository.CreateFakeReportsRepositoryImpl()
}

func CreateReportsController() *controller.ReportController {
	reportsRepository := CreateReportsRepository()
	return controller.CreateReportController(reportsRepository)
}
