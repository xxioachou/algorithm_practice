package main

type pair struct {
	l int
	r int
}

func intersection(a, b pair) pair {
	l := max(a.l, b.l)
	r := min(a.r, b.r)
	if l > r {
		return pair{-1, -1}
	}
	return pair{l, r}
}

type MyCalendarTwo struct {
	seg []pair
	segTwo []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		make([]pair, 0), 
		make([]pair, 0),
	}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	cur := pair{startTime, endTime - 1}
    for _, p := range (*this).segTwo {
		t := intersection(cur, p)
		if t.l != -1 {
			return false
		}
	}	
	for _, p := range (*this).seg {
		t := intersection(cur, p)
		if t.l != -1 {
			(*this).segTwo = append((*this).segTwo, t)
		}
	}
	(*this).seg = append((*this).seg, cur)
	return true
} 

func main() {

}