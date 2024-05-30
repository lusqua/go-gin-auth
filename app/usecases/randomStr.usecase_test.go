package usecases

import (
	"testing"
)

var testedStrings []string

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Test case 1",
			args{32},
		},
		{
			"Test case 2",
			args{32},
		},
		{
			"Test case 3",
			args{32},
		},
		{
			"Test case 4",
			args{32},
		},
		{
			"Test case 5",
			args{32},
		},
		{
			"Test case 6",
			args{32},
		},
		{
			"Test case 7",
			args{32},
		},
		{
			"Test case 8",
			args{32},
		},
		{
			"Test case 9",
			args{32},
		},
		{
			"Test case 10",
			args{32},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				generated := GenerateRandomString(tt.args.n)

				for _, tested := range testedStrings {
					if tested == generated {
						t.Errorf("GenerateRandomString() = %v, already tested", generated)
					}
				}

				testedStrings = append(testedStrings, generated)
			},
		)
	}
}
