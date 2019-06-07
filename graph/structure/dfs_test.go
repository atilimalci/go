package structure

import (
	"log"
	"testing"

	"gotest.tools/assert"
)

func GetGraph() *Graph {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 6)
	g.AddEdge(0, 5)
	g.AddEdge(5, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 10)
	g.AddEdge(9, 12)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 11)
	return g
}

func Test_dfs_PathExist(t *testing.T) {
	g := GetGraph()
	dfs := NewDFS(g, 0)
	nbVertices := 0
	log.Println("Path To 5")
	for x := range dfs.PathTo(5) {
		log.Println(x)
		nbVertices++
	}
	assert.Assert(t, nbVertices == 5, "The number of vertices should be 5")
}

func Test_dfs_NoPath(t *testing.T) {
	g := GetGraph()
	dfs := NewDFS(g, 0)
	nbVertices := 0
	log.Println("Path To 11")
	for x := range dfs.PathTo(11) {
		log.Println(x)
		nbVertices++
	}
	assert.Assert(t, nbVertices == 0, "The number of vertices should be 0")
}
