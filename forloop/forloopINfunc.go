package forloop

import (
	"fmt"
)

func ForloopinFunc() {

	input := []int{100, 20, 30, 40, 50, 60, 200, 80}
	input2 := []float64{100, 20, 30, 40, 50, 60, 200, 80}
	_ = input2

	inputs := [][]int{
		{10, 20, 30},
		{-10, -20, -30},
		{0, -1, 2, -3},
	}

	for i, v := range inputs {
		fmt.Println("inputs", i+1)
		fmt.Println("value =", v)
		fmt.Println("min", v, "=", min(input))
		fmt.Println("min", v, "=", mid(input))
		fmt.Println("max", v, "=", max(input))

	}

	// sqrt := math.Sqrt(float64(input2))
	// fmt.Printf("Result 1: %.1f", sqrt)
	// fmt.Println(sqrt)

}

// input := []int{100, 20, 30, 40, 50, 60, 200, 80}
func mid(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	var value int
	mid := len(input) / 2
	if len(input)%2 == 0 {
		mid = mid - 1
	}

	for i, v := range input {
		if i == mid {
			value = v
		}
	}
	return value
}

func min(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	min := input[0]
	for _, v := range input {
		if v <= min {
			min = v
		}
	}
	return min
}

func max(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	max := input[0]
	for _, v := range input {
		if v >= max {
			max = v
		}
	}
	return max
}

func sudkun5(limit int, sud int) {
	fmt.Println("sudkun", sud)
	fmt.Println("=====")
	for i := 1; i <= limit; i++ {
		fmt.Println(i, "*", sud, "=", i*sud)
	}
}

func average(input []int) int { //ค่าเฉลี่ย
	sum := 0
	for i := range input {
		sum = sum + input[i]
	}
	return sum / len(input)
}

func xbar(list []int) int {
	sumxbar := 0
	for i := range list {
		sumxbar = sumxbar + list[i]
	}
	return sumxbar / len(list)
}

func zixma(list []int) int {
	xi := 0
	sumzixma := 0
	for i := range list {
		xi = (list[i] - xbar(list)) * (list[i] - xbar(list))
		sumzixma = sumzixma + xi
	}
	return sumzixma
}
