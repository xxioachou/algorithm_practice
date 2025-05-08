package main


type OrderedStream struct {
    n 		int
	ptr 	int
	values	[]string
}


func Constructor(n int) OrderedStream {
    return OrderedStream{n: n, values: make([]string, n + 1), ptr: 1}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.values[idKey] = value
	ans := make([]string, 0)
	for this.ptr <= this.n && this.values[this.ptr] != "" {
		ans = append(ans, this.values[this.ptr])
		this.ptr ++
	}
	return ans
}