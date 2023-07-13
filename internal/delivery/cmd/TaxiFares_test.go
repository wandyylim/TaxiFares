package cmd

import (
	mockUC "TaxiFares/internal/mocks/usecase"
	"TaxiFares/model/entity"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestCmdDelivery_validateInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaxiFaresUc := mockUC.NewMockITaxeFares(ctrl)

	type args struct {
		input        string
		prevTime     time.Time
		prevDistance float64
		records      []entity.Record
	}
	tests := []struct {
		name         string
		d            *CmdDelivery
		args         args
		wantTimeVal  time.Time
		wantDistance float64
		wantRecords  []entity.Record
		wantErr      bool
	}{
		{
			name: "tc1 invalid input",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input: "wronginput",
			},
			wantErr: true,
		},
		{
			name: "tc2 error parsing time",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input: "00:00:009.000 0.0",
			},
			wantErr: true,
		},
		{
			name: "tc3 error parsing distance",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input: "00:00:00.000 asd",
			},
			wantTimeVal: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr:     true,
		},
		{
			name: "tc4 invalid time interval",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input:    "00:00:00.000 0.0",
				prevTime: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantTimeVal: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr:     true,
		},
		{
			name: "tc5 interval exceeds limit",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input:    "01:00:00.000 0.0",
				prevTime: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantTimeVal: time.Date(0, 1, 1, 1, 0, 0, 0, time.UTC),
			wantErr:     true,
		},
		{
			name: "tc6 travel distance 0",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input:        "00:01:00.000 0.0",
				prevTime:     time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
				prevDistance: 2.0,
			},
			wantTimeVal: time.Date(0, 1, 1, 0, 1, 0, 0, time.UTC),
			wantErr:     true,
		},
		{
			name: "tc7 sucess",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				input:        "00:01:00.000 5.0",
				prevTime:     time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
				prevDistance: 0.0,
			},
			wantTimeVal:  time.Date(0, 1, 1, 0, 1, 0, 0, time.UTC),
			wantDistance: 5,
			wantRecords: []entity.Record{
				{
					Time:     time.Date(0, 1, 1, 0, 1, 0, 0, time.UTC),
					Distance: 5,
					Diff:     5,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTimeVal, gotDistance, gotRecords, err := tt.d.validateInput(tt.args.input, tt.args.prevTime, tt.args.prevDistance, tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdDelivery.validateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTimeVal, tt.wantTimeVal) {
				t.Errorf("CmdDelivery.validateInput() gotTimeVal = %v, want %v", gotTimeVal, tt.wantTimeVal)
			}
			if gotDistance != tt.wantDistance {
				t.Errorf("CmdDelivery.validateInput() gotDistance = %v, want %v", gotDistance, tt.wantDistance)
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("CmdDelivery.validateInput() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func TestCmdDelivery_processInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaxiFaresUc := mockUC.NewMockITaxeFares(ctrl)

	type args struct {
		records []entity.Record
	}
	tests := []struct {
		name    string
		d       *CmdDelivery
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "tc1 records less than 2",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				records: []entity.Record{},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "tc2 total mileages is 0",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				records: []entity.Record{
					{
						Distance: 0,
						Diff:     0,
					},
					{
						Distance: 0,
						Diff:     0,
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "tc2 total mileages is 0",
			d: &CmdDelivery{
				taxiFaresUc: mockTaxiFaresUc,
			},
			args: args{
				records: []entity.Record{
					{
						Distance: 0,
						Diff:     0,
					},
					{
						Distance: 5,
						Diff:     5,
					},
				},
			},
			wantErr: false,
			mock: func() {
				mockTaxiFaresUc.EXPECT().CalculateFare(gomock.Any()).Return(400.0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if err := tt.d.processInput(tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("CmdDelivery.processInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
