package main

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}
	x := -1
	mul := 0
	if num >= 1000 {
		x = num / 1000
		mul = 1000
	} else if num >= 100 {
		x = num / 100
		mul = 100
	} else if num >= 10 {
		x = num / 10
		mul = 10
	} else {
		x = num
		mul = 1
	}
	if x >= 5 {
		if x == 9 {
			if mul == 100 {
				return "CM" + intToRoman(num - 900)
			}
			if mul == 10 {
				return "XC" + intToRoman(num - 90)
			}
			return "IX" + intToRoman(num - 9)
		}
		if mul == 100 {
			return "D" + intToRoman(num - 500)
		}
		if mul == 10 {
			return "L" + intToRoman(num - 50)
		}
		return "V" + intToRoman(num - 5)
	} 

	if x == 4 {
		if mul == 100 {
			return "CD" + intToRoman(num - 400)
		}
		if mul == 10 {
			return "XL" + intToRoman(num - 40)
		}
		return "IV" + intToRoman(num - 4)
	}
	if mul == 1000 {
		return "M" + intToRoman(num - 1000)
	}
	if mul == 100 {
		return "C" + intToRoman(num - 100)
	}
	if mul == 10 {
		return "X" + intToRoman(num - 10)
	}
	return "I" + intToRoman(num - 1)
}