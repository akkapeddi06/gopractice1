package math1

//This function is used for sum of numbers
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum = sum + v
	}
	return sum
}

//This function is used for sub of numbers
func MySub(xi ...int) int {
	sub := 0
	for _, v := range xi {
		if v > 0 {
			sub = sub + v
		} else {
			sub = sub - v
		}
	}
	return sub
}
