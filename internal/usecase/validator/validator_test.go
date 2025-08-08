package validator

import (
	"testing"

	validatorv10 "github.com/go-playground/validator/v10"
)

type testStruct struct {
	Name string `validate:"non_special_char"`
}

func TestValidateNonSpecialCharacter(t *testing.T) {
	v := validatorv10.New()
	if err := v.RegisterValidation("non_special_char", ValidateNonSpecialCharacter); err != nil {
		t.Fatalf("failed to register validation: %v", err)
	}

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{name: "valid string", value: "Baguette123", valid: true},
		{name: "special characters", value: "Croissant!", valid: false},
		{name: "empty string", value: "", valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(testStruct{Name: tt.value})
			if tt.valid && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !tt.valid && err == nil {
				t.Errorf("expected error but got none")
			}
		})
	}
}
