package maps

import "fmt"

func mapOddEven(arr []int) map[int]string { // รับ[]int{1,2,3} return map  โดยตัวที่เป็นเลขคู่ ="even" , เลขคี่ = "odd"
	m := make(map[int]string)
	for i := range arr {
		v := arr[i]
		if v%2 != 0 { // ถ้า v ซึ่งคือ value ของ arr ที่รับเข้ามา /2 ไม่ลงตัว
			m[v] = "odd" // map `m` key ตัวที่ v จะมี valur = "odd"
			continue     // เหมือน `else if` ถ้าไม่มีcontinue จะกลายเป็นว่า (v ที่ /2 ไม่ลงตัวใน จาก if ที่กำหนดขึ้นมาแล้ว = "odd") จะถูกส่งต่อไปข้างล่างทำให้ต่อ จาก v ที่เข้าเงื่อนไขจะถูกแก้ทับอีกที
		}
		m[v] = "even" // map `m` key ตัวที่ v จะมี valur = "even"

	}
	return m
}

//===========================================================================================//

func MapInFunc() {

	numbers := []int{1, 3, 4, 3, 3, 2, 1}
	countNumbers := count(numbers)
	fmt.Println(countNumbers)

	uniqueNumbers := unique(numbers)
	fmt.Println(uniqueNumbers)

}

// numbers := []int{1, 3, 4, 3, 3, 2, 1} = ["1":2, "2":1 ,"3":3, "4":1 ]
func count(arr []int) map[int]int { // ให้ key ตัวที่ซ้ำ ++
	seen1 := make(map[int]int)
	for _, v := range arr {
		seen1[v]++ //หมายความว่า key ตัวที่ v || value จะ ++ เนื่องจาก default value = 0 || ทำให้ value = 0+1
	} // และ key ของ map จะไม่ถูกส้รางขึ้นซ้ำ || หมายความว่า ถ้าloop เจอตัวเดิม เรากำหนดให้ value ++
	return seen1
}

//============================= เหมือนกัน ==============================

// numbers := []int{1, 1, 3, 4, 3, 3, 2, 1} = {1, 3, 4, 2}
func unique(arr []int) []int { // เก็บเฉพาะข้อมูลที่ไม่ซ้ำ
	box1 := []int{}
	a := make(map[int]bool)
	for _, v := range arr {
		_, ok := a[v] // ท่า _,ok => `ok` คือ bool ซึ่งเอาไว้เช็คว่าใน `map a` มี key v หรือไม่
		if ok {       // หมายความว่าถ้า ถ้ามี key a[v] || a[v] จะเป็น true
			continue
		}

		a[v] = true // ถ้า a[v] = true ให้ append เข้าไปเก็บใน box1
		box1 = append(box1, v)
	}

	return box1
}

func uniqueArt(arr []int) []int {
	result := []int{}
	seen := make(map[int]struct{})

	for _, v := range arr {
		_, ok := seen[v] // ท่า _,ok => `ok` คือ bool ซึ่งเอาไว้เช็คว่าใน `map seen` มี key v หรือไม่
		if ok {
			continue
		}

		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}

func uniqueArr(arr []string) []string {
	result := []string{}
foo:
	for i := range arr {
		v := arr[i]

		for j := range result {
			r := result[j]

			if v == r {
				continue foo
			}
		}

		result = append(result, v)
	}

	return result
}

//===================================================

// arr := []int{-1,2,4,-5}  => {-1:true, 2:false, 4:false, -5:true}
func checkTidlob(n []int) map[int]bool {
	resule := make(map[int]bool)
	for _, v := range n {
		if v < 0 {
			resule[v] = true
			continue
		}
		resule[v] = false
	}
	return resule
}

// arr := []int{-1,2,4,-5}	=>  { true: [-1,-5],false: [2,4] }

func checkTidlob2(n []int) map[bool][]int {
	resule := make(map[bool][]int)
	for _, v := range n {
		key := v < 0
		resule[key] = append(resule[key], v)
	}

	// for _, v := range n {
	// 	if v < 0 {
	// 		resule[v] = append(resule[v], v)
	// 		continue
	// 	}
	// 	resule[v] = append(resule[v], v)
	// }

	return resule

}

//===========================================================
