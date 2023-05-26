package src

import (
	"math"
)

func IsPalindrome(x int64) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	xReversed := reverse(x)

	return xReversed == x
}

func reverse(x int64) int64 {
	sign := "positive"
	if x >= 0 {
		sign = "positive"
	} else {
		sign = "negative"
	}

	x = int64(math.Abs(float64(x)))

	var reversedDigit int64

	for x > 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	if sign == "negative" {
		reversedDigit = reversedDigit * -1
	}

	return reversedDigit
}
