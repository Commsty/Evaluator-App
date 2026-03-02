package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCalculation_Validation(t *testing.T) {
	tests := []struct {
		name        string
		calculation Calculation
		expectError bool
	}{
		{
			name: "valid calculation",
			calculation: Calculation{
				ID:         uuid.NewString(),
				Expression: "2+2",
				Result:     "4",
			},
			expectError: false,
		},
		{
			name: "empty expression",
			calculation: Calculation{
				ID:         uuid.NewString(),
				Expression: "",
				Result:     "4",
			},
			expectError: true,
		},
		{
			name: "empty result",
			calculation: Calculation{
				ID:         uuid.NewString(),
				Expression: "2+2",
				Result:     "",
			},
			expectError: true,
		},
		{
			name: "empty ID",
			calculation: Calculation{
				ID:         "",
				Expression: "2+2",
				Result:     "4",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.calculation.Validate()
			
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCalculationRequest_Validation(t *testing.T) {
	tests := []struct {
		name            string
		calcRequest    CalculationRequest
		expectError     bool
	}{
		{
			name: "valid request",
			calcRequest: CalculationRequest{
				Expression: "2+2",
			},
			expectError: false,
		},
		{
			name: "empty expression",
			calcRequest: CalculationRequest{
				Expression: "",
			},
			expectError: true,
		},
		{
			name: "whitespace only",
			calcRequest: CalculationRequest{
				Expression: "   ",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.calcRequest.Validate()
			
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
