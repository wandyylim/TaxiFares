package taxifares

import "testing"

func TestTaxiFaresUsecase_CalculateFare(t *testing.T) {
	type args struct {
		traveledDistance float64
	}
	tests := []struct {
		name     string
		uc       *TaxiFaresUsecase
		args     args
		wantFare float64
	}{
		{
			name: "case 1 over 1km",
			args: args{
				traveledDistance: 1234,
			},
			wantFare: 560,
		},
		{
			name: "case 2 over 10km",
			args: args{
				traveledDistance: 12340,
			},
			wantFare: 1680,
		},
		{
			name: "case 3 under 1km",
			args: args{
				traveledDistance: 765,
			},
			wantFare: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFare := tt.uc.CalculateFare(tt.args.traveledDistance); gotFare != tt.wantFare {
				t.Errorf("TaxiFaresUsecase.CalculateFare() = %v, want %v", gotFare, tt.wantFare)
			}
		})
	}
}
