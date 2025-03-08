package main

import (
	"reflect"
	"testing"
)

// Табличный тест для функции SortStrings
func TestSortStrings(t *testing.T) {
	type args struct {
		x []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				x: []string{"c", "a", "b"},
			},
			want: []string{"a", "b", "c"},
		},

		{
			name: "Test 2",
			args: args{
				x: []string{"в", "б", "а"},
			},
			want: []string{"а", "б", "в"},
		},

		{
			name: "Test 3",
			args: args{
				x: []string{" ", " ", " "},
			},
			want: []string{" ", " ", " "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortStrings(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
