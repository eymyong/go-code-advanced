package maps

import (
	"fmt"
	"math/rand"
)

type students struct {
	name  string
	class int
}

func newStudents() []students { // สร้าง students
	studentss := make([]students, 10)
	for i := 0; i < 10; i++ {
		studentss[i] = students{
			name:  fmt.Sprintf("students-%d", i+1),
			class: rand.Intn(3) + 1,
		}
	}
	return studentss
}

// จัดกลุ่มนักเรียนใส่ map โดยที่ key จะเป็นห้อง class ||และ value คือนักเรียนที่อยู่ในห้องนั้น (name) ||และ return ออกมา

func arrangeStudents(students []students) map[int][]string {
	m := make(map[int][]string)
	for _, v := range students { // v of students = {students-1 3}
		key := v.class
		valueInkey := m[key]
		valueInkey = append(valueInkey, v.name)
		m[key] = valueInkey
	}
	return m
}

func studentsInRoom(r map[int][]string) map[int]int { // เช็คว่าใน class มี name ทั้งหมดเท่าไหร่
	room := make(map[int]int)
	for i, v := range r {
		room[i] = len(v)
	}
	return room
}

func MapAndstruct() {

	students := newStudents()
	//fmt.Println("students =", students)
	room := arrangeStudents(students)
	fmt.Println(room)

	// numberstudent := studentsInRoom(room)
	// fmt.Println(numberstudent)

	// f := fruitColors{
	// 	"mango":  "yellow",
	// 	"orange": "orange",
	// 	"banana": "yellow",
	// 	"tangmo": "red",
	// }
	// fmt.Println(f)
	// f.addNewFruit("papaya", "orange")
	// fmt.Println(f)

	// new := f.groupFruits()
	// _ = new
	// fmt.Println("map FruitsColor = ", f.groupFruits())

	// fmt.Print(f.groupNumberFruits(), "\n")

}

//=======================================================

type fruitColors map[string]string

func MapAndstruct1() {
	f := fruitColors{
		"mango":  "yellow",
		"orange": "orange",
		"banana": "yellow",
		"tangmo": "red",
	}
	fmt.Println(f)
	f.addNewFruit("papaya", "orange")
	fmt.Println(f)

	new := f.groupFruits()
	_ = new
	fmt.Println("map FruitsColor = ", f.groupFruits())

	fmt.Print(f.groupNumberFruits(), "\n")

	number := []int{1, -1, 4, -5}
	// fmt.Println(checkTidlob(number), "\n")
	// fmt.Println(checkTidlob2(number))
	_ = number

}
func (f fruitColors) addNewFruit(fruit string, color string) { //
	f[fruit] = color
}

func (f fruitColors) getColor(fruitName string) string {
	color := f[fruitName]
	return color
}

func (f fruitColors) groupNumberFruits() map[string]int { //แยกสี โดยแต่ละสีต้องมีผลไม้ที่เป็นสีนั้นๆ
	resule := map[string]int{}
	for fruit := range f {
		color := f[fruit]
		resule[color]++

	}

	return resule
}

func (f fruitColors) groupFruits() map[string][]string { //แยกสี โดยแต่ละสีต้องมีผลไม้ที่เป็นสีนั้นๆ
	result := map[string][]string{}
	for fruit := range f {
		color := f[fruit]
		result[color] = append(result[color], fruit)
	}
	return result
}

func deleteFruit() {

}

// // students := newStudents()
// // fmt.Println(students)

// // arrangeStudents(students)

// // // จัดกลุ่มนักเรียนใส่ map โดยที่ class จะเป็น key value คือนักเรียนที่อยู่ในห้องนั้น และ return ออกมา
// // // ต้องคิดเองด้วยว่าจะ return type อะไร

// // // [{students-1 3} {students-2 2} {students-3 2} {students-4 3} {students-5 2} {students-6 1} {students-7 1} {students-8 1} {students-9 3} {students-10 2}]
// func arrangeStudents(students []students) map[int][]string {
// 	m := make(map[int][]string)

// 	for _, v := range students {
// 		key := v.class
// 		valueInkey := m[key]                  // สร้างvalueInkey ขึ้นมาเพื่อเอาของออกมาจาก map`m[key]`
// 		valueInkey = append(valueInkey, v.id) // พอได้ valueInkey ซึ่งคือของที่เอาออกมาจากmap => ให้เพิ่มของเข้าไปด้วยการ append ก็คือ valueInkey = append(valueInkey, v.id)
// 		m[key] = valueInkey                   // ให้ `m[key]` คือ keyของmap = valueInkey() ซึ่ง `valueInkey` คือของที่มีการเพิ่มเข้าไปแล้ว
// 	}

// 	return m
// }
