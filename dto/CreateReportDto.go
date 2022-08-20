package dto

type CreateReportDto struct {
	MeanVelocity        float32 `json:"mean_velocity" bson:"mean_velocity"`
	MeanAcceleration    float32 `json:"mean_acceleration" bson:"mean_acceleration"`
	Weight              int     `json:"weight" bson:"weight"`
	TrainingMethodology string  `json:"training_methodology" bson:"training_methodology"`
}
