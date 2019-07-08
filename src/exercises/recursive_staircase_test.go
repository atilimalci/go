package exercises

import (
	"testing"
)

func Test_countPaths(t *testing.T) {
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{name: "1 steps", steps: 1, want: 1},
		{name: "3 steps", steps: 3, want: 4},
		{name: "10 steps", steps: 10, want: 274}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPathsRecursive(tt.steps); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPathsRecursiveMemo(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{name: "1 steps", steps: 1, want: 1},
		{name: "3 steps", steps: 3, want: 4},
		{name: "10 steps", steps: 10, want: 274}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPathsRecursiveMemo(tt.steps); got != tt.want {
				t.Errorf("countPathsRecursiveMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPathsDP(t *testing.T) {
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{name: "1 steps", steps: 1, want: 1},
		{name: "3 steps", steps: 3, want: 4},
		{name: "10 steps", steps: 10, want: 274}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPathsDP(tt.steps); got != tt.want {
				t.Errorf("countPathsDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPathsIterative(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "1 steps", args: 1, want: 1},
		{name: "3 steps", args: 3, want: 4},
		{name: "10 steps", args: 10, want: 274}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPathsIterative(tt.args); got != tt.want {
				t.Errorf("countPathsIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
