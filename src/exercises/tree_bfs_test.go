package exercises

import (
	"reflect"
	"testing"
)

func TestNode_bfs(t *testing.T) {
	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test bfs",
			fields: f,
			want:   "12345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.bfs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_bfsRecursive(t *testing.T) {
	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test bfs recursive",
			fields: f,
			want:   "12345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.bfsRecursive(); got != tt.want {
				t.Errorf("Node.bfsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
