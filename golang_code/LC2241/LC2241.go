package main

type ATM struct {
    cnt []int
	val []int
}

func Constructor() ATM {
	return ATM{
		make([]int, 5),
		[]int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int)  {
	for i := range banknotesCount {
		(*this).cnt[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
    tmp := make([]int, 5)
	for i := 4; i >= 0; i-- {
		tmp[i] = min(amount / (*this).val[i], (*this).cnt[i])
		amount -= tmp[i] * (*this).val[i]
	}
	if amount != 0 {
		return []int{-1}
	}
	for i := range tmp {
		(*this).cnt[i] -= tmp[i]
	}
	return tmp
}


func main() {

}