package array

import "fmt"

func ArrayAndMap() {

	// ถ้าเราใส่เลข ตัวเลข ลงใน [] ตอนประกาศตัวแปร array มันจะกำหนด size ตามตัวเลขที่เราใส่
	a := [][2]int{ // กรณีนี้คือ array int ในสุดจะมีสมาชิกได้แค่ 2 ตัว
		{3, 4},
		{8, 9},
		{10, 12},
	}

	mapA := make(map[int]int) //สร้าง map โดยที่ key = a[][0] และ value = a[][1]
	for _, v := range a {     //[3,4],[8,9],[10,12]
		key := v[0]
		value := v[1]
		mapA[key] = value
	}
	fmt.Println("mapA =", mapA) // {3:4,8:9,10:12}

	/*
	   TODO: collect the key from `mapA` to `arrKey` ||and collect the value from `mapA` to `arrVal`
	*/
	arrKey := []int{} // สร้าง []int{} โดยที่ value = key of `mapA`
	for k := range mapA {
		arrKey = append(arrKey, k)
	}
	fmt.Println("arrKey in mapA :", arrKey) // [3,8,10]

	arrVal := []int{} // สร้าง []int{} โดยที่ value = value of `mapA`
	for _, v := range mapA {
		arrVal = append(arrVal, v)
	}
	fmt.Println("arrVal in mapA :", arrVal) // [4,9,12]

	/*
	   TODO: insert key and value from `mapA` into `b` to be the same as `a` (it's fine to incorrect sequence)
	   mapA : map[3:4 8:9 10:12]
	   	a := [][2]int{ {3, 4},{8, 9},{10, 12} }

	*/
	b := [][2]int{}          // สร้าง [][2]int จาก `mapA` โดยที่ value in [2]int{} = {`key`,`value` of `mapA`}
	for i, v := range mapA { //map[3:4 8:9 10:12]
		c := []int{}
		c = append(c, i)
		c = append(c, v)
		b = append(b, [2]int(c))
	}
	fmt.Println("b =", b)

	//===============================================================================================

	/*
	   TODO: use `arrKey` and `arrVal` to insert to `c` to be the same as `a`
	*/
	boxArrKey := arrKey // [3,8,10]
	boxArrVal := arrVal // [4,9,12]

	c := [][2]int{}
	// {3,4}, {8,9}, {10,12}
	for i := range boxArrKey {
		l := boxArrKey[i]
		r := boxArrVal[i]

		c = append(c, [2]int{l, r})
	}
	fmt.Println("c :", c)

	//==============================================================================

	// [][2]int
	// data [][2]int => {1,2}, {3,4}, {5,6}
	var pointers []*[2]int = make([]*[2]int, len(a)) // {p, p, p, p, ...} ; p => [2]int{3,4}
	for i := range a {
		ints2 := a[i]
		pointers[i] = &ints2
	}
	fmt.Println("pointers", pointers)
	/*
		Collect points from a
	*/

}

// 	//================================================================================================

// boxArrKey := arrKey // [3,8,10]
// boxArrVal := arrVal // [4,9,12]
// _ = boxArrVal

// mapNew := make(map[int]int)
// for _, v := range boxArrKey { // key = [3,8,10]
// 	key := v
// 	mapNew[key] = 0
// }
// fmt.Println(mapNew)
// //fmt.Println(mapNew[0])
// mapNew[0] = boxArrVal[0]
// fmt.Println(mapNew)

// for _, v := range boxArrVal { // [4,9,12]
// 	fmt.Println(v)

// }

func mm(a1, a2 []int) map[int]int {
	result := make(map[int]int)
	for _, v := range a1 {
		key := v

		for i := range a2 {
			v := a2[i]
			_, ok := result[key]
			if ok == false {
				result[key] = v
			}
		}
	}
	return result
}
