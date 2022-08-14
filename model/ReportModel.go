package model

type ReportModel struct {
	Id                  string  `json:"id"`
	MeanVelocity        float32 `json:"mean_velocity"`
	MeanAcceleration    float32 `json:"mean_acceleration"`
	Weight              int     `json:"weight"`
	TrainingMethodology string  `json:"training_methodology"`
}

const (
	VBT          = "VBT"
	FreeTraining = "Treino Livre"
)
