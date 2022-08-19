package repository

import (
	"errors"
	"fmt"
	"poc-push-app-api/domain/model"
)

type FakeReportsRepositoryImpl struct {
	reports []model.ReportModel
}

func CreateFakeReportsRepositoryImpl() *FakeReportsRepositoryImpl {
	return &FakeReportsRepositoryImpl{
		reports: []model.ReportModel{
			{
				Id:                  "2313445456",
				MeanVelocity:        0.65,
				MeanAcceleration:    0.25,
				Weight:              50,
				TrainingMethodology: model.VBT,
			},
			{
				Id:                  "534532343e3",
				MeanVelocity:        0.68,
				MeanAcceleration:    0.32,
				Weight:              40,
				TrainingMethodology: model.FreeTraining,
			},
		},
	}
}

func (fakeReportsRepository FakeReportsRepositoryImpl) GetAll() ([]model.ReportModel, error) {
	return fakeReportsRepository.reports, nil
}

func (fakeReportsRepository *FakeReportsRepositoryImpl) GetById(id string) (model.ReportModel, error) {

	foundedReport := model.ReportModel{}

	for _, report := range fakeReportsRepository.reports {
		if report.Id == id {
			foundedReport = report
		}
	}

	if (model.ReportModel{} == foundedReport) {
		return model.ReportModel{}, errors.New(fmt.Sprintf("O Report com o id %s não foi localizado", id))
	}

	return foundedReport, nil
}

func (fakeReportsRepository *FakeReportsRepositoryImpl) Create(report model.ReportModel) (model.ReportModel, error) {
	fakeReportsRepository.reports = append(fakeReportsRepository.reports, report)
	return report, nil
}

func (fakeReportsRepository *FakeReportsRepositoryImpl) Delete(id string) error {

	reports := fakeReportsRepository.reports

	for index, report := range reports {
		if report.Id == id {
			fakeReportsRepository.reports = append(reports[:index], reports[index+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Não foi localizado o report de id %s", id))
}
