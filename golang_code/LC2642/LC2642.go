package main

const inf = 1e15

type Graph struct {
	n int
    d [][]int
}


func Constructor(n int, edges [][]int) Graph {
    var g Graph 
	g.n = n
	g.d = make([][]int, n)
	for i := 0; i < n; i++ {
		g.d[i] = make([]int, n)
		g.d[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.d[i][j], g.d[j][i] = inf, inf
		}
	}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.d[u][v] = min(g.d[u][v], w)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g.d[i][j] = min(g.d[i][j], g.d[i][k] + g.d[k][j])
			}
		}
	}
	return g
}


func (this *Graph) AddEdge(edge []int)  {
    u, v, w := edge[0], edge[1], edge[2]	
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			this.d[i][j] = min(this.d[i][j], this.d[i][u] + this.d[v][j] + w)
		}
	}
}


func (this *Graph) ShortestPath(node1 int, node2 int) int {
    ans := this.d[node1][node2]
	if ans >= inf {
		ans = -1
	}
	return ans
}

func main() {

}
