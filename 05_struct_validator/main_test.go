package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidateStruct(t *testing.T) {
	tests := []struct {
		name    string
		input   AA
		wantErr bool
	}{
		{
			name: "All values nil",
			input: AA{
				Status: "ok",
				Value1: nil,
				Value2: nil,
				Value3: nil,
			},
			wantErr: true,
		},
		{
			name: "All values within range",
			input: AA{
				Status: "ok",
				Value1: float64Ptr(50),
				Value2: float64Ptr(50),
				Value3: float64Ptr(50),
			},
			wantErr: false,
		},
		{
			name: "Value1 out of range",
			input: AA{
				Status: "ok",
				Value1: float64Ptr(150),
				Value2: float64Ptr(50),
				Value3: float64Ptr(50),
			},
			wantErr: true,
		},
		{
			name: "Value2 out of range",
			input: AA{
				Status: "ok",
				Value1: float64Ptr(50),
				Value2: float64Ptr(150),
				Value3: float64Ptr(50),
			},
			wantErr: true,
		},
		{
			name: "Value3 out of range",
			input: AA{
				Status: "ok",
				Value1: float64Ptr(50),
				Value2: float64Ptr(50),
				Value3: float64Ptr(150),
			},
			wantErr: true,
		},
		{
			name: "Status not ok",
			input: AA{
				Status: "not_ok",
				Value1: nil,
				Value2: nil,
				Value3: nil,
			},
			wantErr: false,
		},
	}

	v := validator.New()
	v.RegisterStructValidation(validateStruct, AA{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}
