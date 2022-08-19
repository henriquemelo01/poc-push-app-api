package di

import (
	"poc-push-app-api/controller"
	"poc-push-app-api/services/repository"
)

func CreateMongoReportsRepositoryImpl() *repository.MongoReportsRepositoryImpl {
	return repository.CreateMongoReportsRepositoryImpl()
}

func CreateFakeReportsRepositoryImpl() *repository.FakeReportsRepositoryImpl {
	return repository.CreateFakeReportsRepositoryImpl()
}

func CreateReportsController() *controller.ReportController {
	reportsRepository := CreateMongoReportsRepositoryImpl()
	return controller.CreateReportController(reportsRepository)
}
