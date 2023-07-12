package taxifares

import (
	constant "TaxiFares/model/constant"
	"math"
)

func (uc *TaxiFaresUsecase) CalculateFare(traveledDistance float64) (fare float64) {

	if traveledDistance > 0 {
		distance := traveledDistance

		fare = uc.calculateFare(distance)
	}

	return fare

}

func (uc *TaxiFaresUsecase) calculateFare(distance float64) float64 {
	baseFare := constant.BaseFare

	// if distance less than 1km
	if distance <= 1000.0 {
		return baseFare
	} else if distance <= 10000.0 {
		// if distance more than 1km but less than 10km
		additionalCharge := math.Ceil(distance/constant.Distance1) * constant.DistanceRate1
		return baseFare + additionalCharge
	} else {
		// over 10km
		additionalCharge := (10000.0 / constant.Distance1) * constant.DistanceRate1                   // Additional charge up to 10 km
		additionalCharge += math.Ceil((distance-10000.0)/constant.Distance2) * constant.DistanceRate2 // Additional charge beyond 10 km
		return baseFare + additionalCharge
	}
}
