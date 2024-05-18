package forloop

import "fmt"

func Loop() {

}

func Pattern1() {

	for i := 1; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println("i=", i, "j=", j)
		}
		fmt.Println("===")
		fmt.Println(i)
	}
}

func Pattern2() {
	for i := 1; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println("aa")
	}
}

func Pattern3() {

	for i := 1; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("%d%d", i, j)
		}
		fmt.Println("finish loop j")
	}
}

func loop() {
	var input string
	for {
		fmt.Scanf("%s", &input) // รับ input `string` ที่พิมพ์เข้าไปใน termonal หลังจาก run func
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}

func loopcars() {
	cars := []string{"mazda", "toyota", "honda"}
	for i := 0; i <= len(cars); i++ {
		fmt.Println(cars[i])
	}
	for i, c := range cars {
		fmt.Println(c, cars[i])
	}

}

func indexInt() []int { // สร้างตัวเลขด้วยการ loop => add in []int{}
	var box []int
	for i := 1; i <= 20; i++ {
		box = append(box, i)
	}
	return box
}

//[3 6 9 12 15 18 21 24 27 30]

func loop1() []int { // ไม่รับค่า,loop next ไปจนถึง i ที่กำหนด โดยกำหนดเงื่อนไข
	box := []int{}
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			box = append(box, i)
		}
	}
	return box
}

func loop2(limit int) []int { // รับค่า limit int ,loop next to limit
	var box []int
	for i := 1; i <= limit; i++ {
		box = append(box, i)
	}
	return box
}

func loop3(input1 []int) []int { // รับ input1 []int ,loop next ,loop in input, ใช้ range(1)index,2)value)
	box := []int{}
	for _, v := range input1 {
		box = append(box, v)
	}
	return box
}

func loop4(input1 []int) []int { // รับ input1 []int, loop back
	box := []int{}
	for i := len(input1) - 1; i >= 0; i-- {
		v := input1[i]
		box = append(box, v)
	}
	return box
}

func loop5(input1 []int) []int { // รับ input1 []int ,loop next ,loop in input ถึง mid(input1)
	mid := len(input1) / 2
	if len(input1)%2 == 0 {
		mid = mid - 1
	}

	box := []int{}
	for i := range input1 {
		v := input1[i]
		box = append(box, v)
		if i == mid {
			break
		}
	}
	return box
}

func loop6(input1 []int) []int { // รับ input1 []int ,loop back ,loop in input ถึง mid(input1)
	mid := len(input1) / 2
	if len(input1)%2 == 0 {
		mid = mid - 1
	}

	box := []int{}
	for i := len(input1) - 1; i >= 0; i-- {
		v := input1[i]
		box = append(box, v)
		if i == mid+1 {
			break
		}
	}
	return box
}

func loop7(input1 []int, index int) []int { // รับ input1 []int, index int , loop input1 * index, return []
	box := []int{}
	for _, v := range input1 {
		vv := v * index
		box = append(box, vv)
	}
	return box
}

func loop8(input1 []int, index int) []int { // รับ input1 []int, index int , loop input1 * index, add if,return []
	box := []int{}
	for _, v := range input1 {
		v = v * index
		if v%4 == 0 {
			box = append(box, v)
		}
	}
	return box
}

/*  การบ้าน
1. average find mean from array of int
	- input = [1,2,3] output = 2
2. find duplicate numbers from two arrays
	-  if input1 = [1,2,3] input2 = [2,3,7] this function will return [2,3]
3. find prime number that less than limit
	- func will return array of prime number that less than limit
	- input = 7 output =[2,3,5]

	[1,2,3,4]
	[5,2,9,1,8,32,246462,4]
	[]

*/
