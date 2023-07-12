package usecase

type ITaxeFares interface {
	CalculateFare(traveledDistance float64) float64
}
