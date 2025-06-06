package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	const inf = 0x3f3f3f3f
	n := len(edges)
	dis1 := make([]int, n)
	dis2 := make([]int, n)
	for i := range dis1 {
		dis1[i], dis2[i] = inf, inf
	}
	calc_dis := func(st int, dis []int) {
		for d, x := 0, st; x != -1 && dis[x] == inf; x = edges[x] {
			dis[x] = d
			d++
		}
	}
	calc_dis(node1, dis1)
	calc_dis(node2, dis2)
	ans, idx := inf, -1
	for i := range dis1 {
		if ans > max(dis1[i], dis2[i]) {
			ans = max(dis1[i], dis2[i])
			idx = i
		}
	}
	return idx
}

func main() {

}
