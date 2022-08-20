package repository

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		reportDtoId := reportDTO.Id.(primitive.ObjectID).Hex()

		reportModel := model.ReportModel{
			Id:                  reportDtoId,
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

func (mr *MongoReportsRepositoryImpl) Create(createReportDto dto.CreateReportDto) (model.ReportModel, error) {

	client, connectionError := mr.connectToMongoClient()

	if connectionError != nil {
		return model.ReportModel{}, connectionError
	}

	ctx := context.Background()

	createdReport, createReportDocumentsError := client.Database("PushApp").Collection("reports").InsertOne(ctx, createReportDto)
	if createReportDocumentsError != nil {
		return model.ReportModel{}, createReportDocumentsError
	}

	createdReportId := createdReport.InsertedID.(primitive.ObjectID).Hex()

	report := model.ReportModel{
		Id:                  createdReportId,
		MeanVelocity:        createReportDto.MeanVelocity,
		MeanAcceleration:    createReportDto.MeanAcceleration,
		Weight:              createReportDto.Weight,
		TrainingMethodology: createReportDto.TrainingMethodology,
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

	ctx := context.Background()

	objectId, createObjectIdError := primitive.ObjectIDFromHex(id)
	if createObjectIdError != nil {
		return model.ReportModel{}, createObjectIdError
	}

	searchReportByIdQuery := bson.D{{"_id", objectId}}

	var reportDTO dto.ReportDTO
	getReportDocumentsError :=
		client.Database("PushApp").Collection("reports").FindOne(ctx, searchReportByIdQuery).Decode(&reportDTO)

	if getReportDocumentsError != nil {
		return model.ReportModel{}, getReportDocumentsError
	}

	reportDtoId := reportDTO.Id.(primitive.ObjectID).Hex()

	reportModel := model.ReportModel{
		Id:                  reportDtoId,
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

	objectId, createObjectIdError := primitive.ObjectIDFromHex(id)
	if createObjectIdError != nil {
		return createObjectIdError
	}

	deleteReportByIdQuery := bson.D{{"_id", objectId}}

	if connectionError != nil {
		return connectionError
	}

	_, deletedReportError := client.Database("PushApp").Collection("reports").DeleteOne(ctx, deleteReportByIdQuery)
	if deletedReportError != nil {
		return deletedReportError
	}

	if disconnectToClientError := client.Disconnect(context.TODO()); disconnectToClientError != nil {
		return disconnectToClientError
	}

	return nil
}
