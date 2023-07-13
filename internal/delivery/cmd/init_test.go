package cmd

import (
	"TaxiFares/internal/usecase"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		taxiFaresUc usecase.ITaxeFares
	}
	tests := []struct {
		name    string
		args    args
		want    *CmdDelivery
		wantErr bool
	}{
		{
			name: "tc 1 success",
			want: &CmdDelivery{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.taxiFaresUc)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
