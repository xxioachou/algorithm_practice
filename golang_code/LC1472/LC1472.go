package main

type BrowserHistory struct {
    index 	int
	sites	[]string
}


func Constructor(homepage string) BrowserHistory {
	t := BrowserHistory{
		index: 0,
		sites: make([]string, 0),
	}
	t.sites = append(t.sites, homepage)
	return t
}


func (this *BrowserHistory) Visit(url string)  {
	this.sites = this.sites[:this.index + 1]
	this.sites = append(this.sites, url)
	this.index ++
}


func (this *BrowserHistory) Back(steps int) string {
    steps = min(steps, this.index)
	this.index -= steps
	return this.sites[this.index]
}


func (this *BrowserHistory) Forward(steps int) string {
	steps = min(steps, len(this.sites) - 1 - this.index)
	this.index += steps
	return this.sites[this.index]
}