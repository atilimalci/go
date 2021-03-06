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

//	    1
//	  /   \
//   2     3
// /   \
//4     5
//
//42513
func TestNode_dfsInorder(t *testing.T) {
	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test dfsInorder",
			fields: f,
			want:   "42513"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.dfsInorder(); got != tt.want {
				t.Errorf("Node.dfsInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

//	    1
//	  /   \
//   2     3
// /   \
//4     5
//
//42513
func TestNode_dfsInorderRecursive(t *testing.T) {
	f := getTestData()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test dfsInorderRecursive",
			fields: f,
			want:   "42513"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &Node{
				i:     tt.fields.i,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := root.dfsInorderRecursive(); got != tt.want {
				t.Errorf("Node.dfsInorderRecursive() = %v, want %v", got, tt.want)
			}
		})
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

func TestNode_dfsPreorderRecursive(t *testing.T) {

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
			if got := root.dfsPreorderRecursive(); got != tt.want {
				t.Errorf("Node.dfsPreorderRecursive() = %v, want %v", got, tt.want)
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
