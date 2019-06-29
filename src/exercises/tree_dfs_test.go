package exercises

import (
	"reflect"
	"testing"
)

type fields struct {
	i     int
	left  *Node
	right *Node
}

func getTestData() fields {
	n0 := &Node{4, nil, nil}
	n1 := &Node{5, nil, nil}
	n2 := &Node{2, n0, n1}
	n3 := &Node{3, nil, nil}

	return fields{i: 1, left: n2, right: n3}
}

func Test_dfsInorder(t *testing.T) {
	n0 := &Node{4, nil, nil}
	n1 := &Node{5, nil, nil}
	n2 := &Node{2, n0, n1}
	n3 := &Node{3, nil, nil}
	n4 := &Node{1, n2, n3}

	s := n4.dfsInorder()
	if s != "42513" {
		t.Errorf("Expected 42513 but got %s", s)
	}
}

//	    1
//	  /   \
//   2     3
// /   \
//4     5
func TestNode_dfsPreorder(t *testing.T) {

	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test dfs Preorder",
			fields: f,
			want:   "12453"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.dfsPreorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.dfsPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

//	    1
//	  /   \
//   2     3
// /   \
//4     5

// 4 5 2 3 1

func TestNode_dfsPostorder(t *testing.T) {
	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test dfs Postorder",
			fields: f,
			want:   "45231"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.dfsPostorder(); got != tt.want {
				t.Errorf("Node.dfsPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_dfsPosterderRecursive(t *testing.T) {
	f := getTestData()
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test dfsPosterderRecursive",
			fields: f,
			want:   "45231"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.dfsPosterderRecursive(); got != tt.want {
				t.Errorf("Node.dfsPosterderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
