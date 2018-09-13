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
