package usecase

type ITaxeFares interface {
	CalculateFare(elapsedTime, traveledDistance float64) int
}
