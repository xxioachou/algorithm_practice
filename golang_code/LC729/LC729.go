package main

type pair struct {
	l int
	r int
}

func intersection(a, b pair) bool {
	l := max(a.l, b.l)
	r := min(a.r, b.r)
	return l <= r
}

type MyCalendar struct {
	seg []pair
}

func Constructor() MyCalendar {
	return MyCalendar{make([]pair, 0)}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	tmp := pair{startTime, endTime - 1}
    for _, p := range (*this).seg {
		if intersection(tmp, p) {
			return false
		}
	}	
	(*this).seg = append((*this).seg, tmp)
	return true
} 

func main() {

}