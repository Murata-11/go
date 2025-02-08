package main

import (
	"reflect"
	"testing"
)

func TestIntToBinaryArray1Based(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  []string
	}{
		{
			name:  "Binary of 0",
			input: 0,
			want:  []string{"", "[1] 0"},
		},
		{
			name:  "Binary of 1",
			input: 1,
			want:  []string{"", "[1] 1"},
		},
		{
			name:  "Binary of 2",
			input: 2,
			want:  []string{"", "[1] 1", "[2] 0"},
		},
		{
			name:  "Binary of 10",
			input: 10,
			want:  []string{"", "[1] 1", "[2] 0", "[3] 1", "[4] 0"},
		},
		{
			name:  "Binary of 127",
			input: 127,
			want:  []string{"", "[1] 1", "[2] 1", "[3] 1", "[4] 1", "[5] 1", "[6] 1", "[7] 1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToBinaryArray1Based(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToBinaryArray1Based(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
