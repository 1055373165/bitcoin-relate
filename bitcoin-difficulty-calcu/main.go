package main

import (
	"fmt"
	"math"
)

func main() {
	attack_q := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	left_block := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("%2s \t %s \t %13s\n\n", "q", "z", "probility")
	for _, q := range attack_q {
		for _, z := range left_block {
			probility := AttackSuccessProbability(q, z)
			fmt.Printf("%v \t %v \t %v\n", q, z, probility)
		}
		fmt.Println()
		fmt.Println("_________________________________________")
		fmt.Println()
	}
}

func AttackSuccessProbability(q float64, z int) float64 {
	// p is the probability of honester
	var p float64 = 1.0 - q
	var lambda float64 = float64(z) * (q / p)
	var sum float64 = 1.0

	var i, k int

	for ; k <= z; k++ {
		var possion float64 = math.Exp(-lambda)
		// 模拟 / k!  k! 从 1 * ... * k
		for i = 1; i <= k; i++ {
			possion *= lambda / float64(i)
		}
		sum -= possion * (1 - math.Pow(q/p, float64(z-k)))
	}

	return sum
}
