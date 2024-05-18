package homework1

import "fmt"

func HomeworkArray1() {
	number := []int{1, 1, 2, 2, 3, 3}

	boxCherkDouble := checkDouble(number)
	fmt.Println(boxCherkDouble)

	// boxCount := count(number)
	// fmt.Println(boxCount)

}

// number := []int{1,1,2,2,3,3} = {1,2,3}
func checkDouble(arr []int) []int { // เก็บข้อมูลที่ไม่ซ้ำ
	result := []int{}
	m := make(map[int]bool)
	/*
		เราต้องการเก็บข้อมูลใน `arr` ที่ไม่เอาตัวซ้ำ => จึงสร้าง`map`ขึ้นมาโดยที่ข้อมูลที่รับเข้ามาเป็น key (เพราะkeyในmap จะไม่สามารถซ้ำกันได้)
		โดยให้ value in map เป็น `bool` เพราะเราจะเช็คว่ามี key ใน map แล้วหรือยัง
	*/
	for _, v := range arr {
		_, ok := m[v]   // (*ท่านี้เป็นท่าที่ใช้กับ `map`) // โดยที่ ok มี type เป็น bool เสมอ //( _, ok := m[v] ) เป็นการเช็คว่า ok มีkey `m[v]` หรือไม่ // โดยที่ถ้ามีเป็น true ,ถ้าไม่มีเป็น false
		if ok == true { // ถ้า ok == true(หรือหมายความว่าถ้ามี key `ok` แล้ว) ให้เข้ามาทำใน if โดยการที่ให้ `continue` คือการออกจาก loop
			continue
		}
		m[v] = true // (ทำต่อจาก line22 คือต้องไม่เข้าเงื่อนไข line 24) หมายความว่ายังไม่มี key `m[v]` => จึงให้ `m[v]` = true
		result = append(result, v)
	}
	return result
}

// numbers := []int{1,1,2,2,3,3} = ["1":2, "2":2 ,"3":2]
func count(arr []int) map[int]int { // เก็บข้อมูลตัวที่ซ้ำ
	result := make(map[int]int)
	for _, v := range arr {
		result[v]++
	}

	return result

}
