package main

func mian() {

}

//860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	amt05, amt10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			amt05++
		} else if bill == 10 {
			if amt05 == 0 {
				return false
			}
			amt05--
			amt10++
		} else {
			if amt05 > 0 && amt10 > 0 {
				amt05--
				amt10--
			} else if amt05 >= 3 {
				amt05 -= 3
			} else {
				return false
			}
		}
	}
	return true

}

//121. 买卖股票的最佳时机
func maxProfit02(prices []int) int {
	a := 0
	for i := 0; i < len(prices); i++ {
		c := prices[i] - prices[i-1]
		if c < 0 {
			c = 0
		}
		if c > a {
			a = c
		}
	}
	return a
}
