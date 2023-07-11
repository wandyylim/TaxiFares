package taxifares

import (
	constant "TaxiFares/model/constant"
)

func (uc *TaxiFaresUsecase) CalculateFare(elapsedTime, traveledDistance float64) int {
	fare := constant.BaseFare

	if traveledDistance > 0 {
		distance := traveledDistance

		if distance <= constant.Distance1 {
			fare += int(distance / constant.Distance1 * constant.DistanceRate1)
		} else {
			fare += int(constant.Distance1 / constant.Distance1 * constant.DistanceRate1)

			distance -= constant.Distance1
			fare += int(distance / constant.Distance2 * constant.DistanceRate2)
		}
	}

	// Fare per minute calculation
	farePerMinute := float64(fare) / elapsedTime
	farePerMinute = uc.round(farePerMinute, 0.01)

	// Adding the minimum fare per minute
	minimumFarePerMinute := float64(constant.BaseFare) / 60
	farePerMinute += minimumFarePerMinute

	return int(farePerMinute * elapsedTime)

}

func (uc *TaxiFaresUsecase) round(value, unit float64) float64 {
	return float64(int((value/unit)+0.5)) * unit
}
