package structure

type Graph struct {
	Adjacencies map[int][]int
}

func NewGraph() *Graph {
	g := &Graph{}
	g.Adjacencies = make(map[int][]int)
	return g
}

func (g *Graph) AddEdge(v, w int) {
	if _, ok := g.Adjacencies[v]; !ok {
		g.Adjacencies[v] = make([]int, 0)
	}
	g.Adjacencies[v] = append(g.Adjacencies[v], w)

	if _, ok := g.Adjacencies[w]; !ok {
		g.Adjacencies[w] = make([]int, 0)
	}
	g.Adjacencies[w] = append(g.Adjacencies[w], v)

}

func (g *Graph) Adj(v int) <-chan int {
	ch := make(chan int)
	go func() {
		if adj, ok := g.Adjacencies[v]; ok {
			for val := range iterateArray(adj) {
				ch <- val
			}
		}
		close(ch)
	}()
	return ch
}

func iterateArray(bag []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range bag {
			ch <- val
		}
		close(ch)
	}()
	return ch
}
