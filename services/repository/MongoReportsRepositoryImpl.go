package repository

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"poc-push-app-api/domain/model"
	"poc-push-app-api/dto"
	"time"
)

type MongoReportsRepositoryImpl struct{}

func CreateMongoReportsRepositoryImpl() *MongoReportsRepositoryImpl {
	return &MongoReportsRepositoryImpl{}
}

func (mr *MongoReportsRepositoryImpl) connectToMongoClient() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if loadingDotEnvFileError := godotenv.Load(".env"); loadingDotEnvFileError != nil {
		return nil, loadingDotEnvFileError
	}

	client, clientConnectionError := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if clientConnectionError != nil {
		return nil, clientConnectionError
	}

	return client, nil
}

func (mr *MongoReportsRepositoryImpl) GetAll() ([]model.ReportModel, error) {

	client, connectionError := mr.connectToMongoClient()

	if connectionError != nil {
		return nil, connectionError
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	ctx := context.Background()

	cursor, getReportsDocumentsError := client.Database("PushApp").Collection("reports").Find(ctx, bson.M{})

	if getReportsDocumentsError != nil {
		return []model.ReportModel{}, getReportsDocumentsError
	}

	var reportsDTO []dto.ReportDTO
	if parseDocError := cursor.All(ctx, &reportsDTO); parseDocError != nil {
		return nil, parseDocError
	}

	var reportsModel []model.ReportModel
	for _, reportDTO := range reportsDTO {

		reportModel := model.ReportModel{
			Id:                  reportDTO.Id,
			MeanVelocity:        reportDTO.MeanVelocity,
			MeanAcceleration:    reportDTO.MeanAcceleration,
			Weight:              reportDTO.Weight,
			TrainingMethodology: reportDTO.TrainingMethodology,
		}

		reportsModel = append(reportsModel, reportModel)
	}

	if disconnectToClientError := client.Disconnect(context.TODO()); disconnectToClientError != nil {
		return nil, disconnectToClientError
	}

	return reportsModel, nil
}

func (mr *MongoReportsRepositoryImpl) Create(report model.ReportModel) (model.ReportModel, error) {

	client, connectionError := mr.connectToMongoClient()

	if connectionError != nil {
		return model.ReportModel{}, connectionError
	}

	ctx := context.Background()

	reportDTO := dto.ReportDTO{
		Id:                  report.Id,
		MeanVelocity:        report.MeanVelocity,
		MeanAcceleration:    report.MeanAcceleration,
		Weight:              report.Weight,
		TrainingMethodology: report.TrainingMethodology,
	}

	_, createReportDocumentsError := client.Database("PushApp").Collection("reports").InsertOne(ctx, reportDTO)
	if createReportDocumentsError != nil {
		return model.ReportModel{}, createReportDocumentsError
	}

	if disconnectToClientError := client.Disconnect(context.TODO()); disconnectToClientError != nil {
		return model.ReportModel{}, disconnectToClientError
	}

	return report, nil
}

func (mr *MongoReportsRepositoryImpl) GetById(id string) (model.ReportModel, error) {

	client, connectionError := mr.connectToMongoClient()

	if connectionError != nil {
		return model.ReportModel{}, connectionError
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	ctx := context.Background()

	var reportDTO dto.ReportDTO
	getReportDocumentsError := client.Database("PushApp").Collection("reports").FindOne(ctx, bson.D{{"id", id}}).Decode(&reportDTO)

	if getReportDocumentsError != nil {
		return model.ReportModel{}, getReportDocumentsError
	}

	reportModel := model.ReportModel{
		Id:                  reportDTO.Id,
		MeanVelocity:        reportDTO.MeanVelocity,
		MeanAcceleration:    reportDTO.MeanAcceleration,
		Weight:              reportDTO.Weight,
		TrainingMethodology: reportDTO.TrainingMethodology,
	}

	if disconnectToClientError := client.Disconnect(context.TODO()); disconnectToClientError != nil {
		return model.ReportModel{}, disconnectToClientError
	}

	return reportModel, nil
}

func (mr *MongoReportsRepositoryImpl) Delete(id string) error {

	client, connectionError := mr.connectToMongoClient()

	ctx := context.Background()

	if connectionError != nil {
		return connectionError
	}

	_, deletedReportError := client.Database("PushApp").Collection("reports").DeleteOne(ctx, bson.D{{"id", id}})
	if deletedReportError != nil {
		return deletedReportError
	}

	if disconnectToClientError := client.Disconnect(context.TODO()); disconnectToClientError != nil {
		return disconnectToClientError
	}

	return nil
}
