package array

import "fmt"

func Array1() {

	inputs := [][]int{
		{10, 20, 30},
		{-10, -20, -30},
		{0, -1, 2, -3},
	}

	filterEven(inputs)
	filterEven2(inputs)

	input := []int{100, 20, 30, 40, 50, 60, 200, 80}
	n := 10
	aa := listPercentOf(input, n)

	for i := 0; i < len(aa); i++ {
		fmt.Println("===ผลสอบ===")
		fmt.Println("คะแนนที่ได้ =", input[i], ": คะแนนเต็ม =", n)
		//fmt.Println("คุณสอบได้ =", percentOf(input[i], n), "%")
		fmt.Println("คุณสอบได้", aa[i], "%")
		if input[i] == n {
			fmt.Println("เก่งมาก")
		} else if input[i] <= 3 {
			fmt.Println("คุณสอบตก")
		}

	}

}

// data   = [[1,2,3], [4,5,6]]
// output = [2,4,6]
func filterEven(data1 [][]int) []int {
	result := []int{}

	for _, v := range data1 {
		for _, v2 := range v {
			if v2%2 == 0 {
				result = append(result, v2)
			}
		}
	}
	return result

}

// data   = [[1,2,3], [4,5,6], [1,3,5], [10,20]]
// output = [[2], [4,6], [], [10,20]]
func filterEven2(data2 [][]int) [][]int {
	result := [][]int{}
	for _, v := range data2 {
		box := []int{}
		for _, v2 := range v {
			if v2%2 == 0 {
				box = append(box, v2)
			}
		}
		result = append(result, box)
	}
	return result
}

func percentOf(part int, total int) float64 { //หา%
	fmt.Println("===ผลสอบ===")
	fmt.Println("คะแนนที่ได้ =", part, ": คะแนนเต็ม =", total)
	fmt.Println("คุณสอบได้ =", (float64(part)*float64(100))/float64(total), "%")
	return (float64(part) * float64(100)) / float64(total)
}

func listPercentOf(points []int, total int) []float64 { //หา % in array[]
	result := make([]float64, len(points))
	for i := range points {
		v := points[i]
		result[i] = percentOf(v, total)
	}
	return result
}
