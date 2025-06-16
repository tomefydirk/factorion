package main

import (
	"fmt"
)

//145!=1!+4!+5! factorion
//limit 7 chiffres pourquoi (mba tediavo)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func isFactorion(n int) bool {
	sum := 0
	temp := n
	for temp > 0 {
		digit := temp % 10
		sum += factorial(digit)
		temp /= 10
	}
	return sum == n
}

func main() {
	fmt.Println("Nombres factorions en base 10:")
	for i := 1; i <= 1000000; i++ {
		if isFactorion(i) {
			fmt.Println(i)
		}
	}
}
