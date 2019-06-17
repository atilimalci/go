package structure

import "github.com/golang-collections/collections/stack"

type DFS struct {
	Source int
	Marked map[int]bool
	Adj    map[int]int
	G      *Graph
}

func NewDFS(g *Graph, source int) *DFS {
	dfs := &DFS{
		Marked: make(map[int]bool),
		Adj:    make(map[int]int),
		G:      g,
		Source: source,
	}
	dfs.dfs(source)
	return dfs
}

func (g *DFS) dfs(v int) {
	g.Marked[v] = true
	for w := range g.G.Adj(v) {
		if !g.Marked[w] {
			g.dfs(w)
			g.Adj[w] = v
		}
	}
}

func (g *DFS) HasPathTo(v int) bool {
	return g.Marked[v]
}

func (g *DFS) PathTo(v int) <-chan interface{} {
	ts := stack.New()
	if g.HasPathTo(v) {

		for x := v; x != g.Source; x = g.Adj[x] {
			ts.Push(x)
		}
		ts.Push(g.Source)
	}
	return iterate(ts)
}

func iterate(s *stack.Stack) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for {
			if s.Len() == 0 {
				break
			}
			ch <- s.Pop()
		}
		close(ch)

	}()
	return ch
}
