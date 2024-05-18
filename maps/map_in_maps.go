package maps

import "fmt"

func Mapinmap() {

	//m := make(map[string]map[int][]string)

	// for k1 := range m {
	// 	for k2, v := range m[k1] {
	// 		_, ok = m[k1]
	// 	}
	// }

	// i := m["hee"]

	// g := i[7]
	// h := m["hee"][7]
	// e := m["hee"][8][8]
	// fmt.Println(g, h, e)

	a := map[string]map[int][]string{
		"yong": {
			1: {"kuy"},
			4: {"foo", "bar"},
		},

		"noon": {
			0: {"lol", "lmao"},
			//2: {"q", "qq"},
		},
	}
	_ = a

	// เช็คได้แบบนี้
	// a["noon"][0][1] == "lmao"
	// len(m["yong"]) == 2
	fmt.Println(len(a["noon"]))
	fmt.Println(len(a["yong"][4]))
	fmt.Println(a["yong"][4][0])

	m := map[string]map[int][]string{
		"yong": {
			1: {"y", "yy"},
			2: {"o", "oo"},
		},
		// "noOn": {
		// 	11: {"n", "nn"},
		// },
	}

	fmt.Println("Map m =\n", m)
	//================  add m ================= เพิ่ม key ลงใน map `m` ซึ่งใน `m` = map[string]map[int][]string
	// z := map[int][]string{}    // สร้าง `z` ย่อย value ของ m ขึ้นมาก่อน เพราะ value ของ `m` คือ  map[int][]string
	// z[55] = []string{"p", "a"} // add
	// z[66] = []string{"pp", "aa"}
	// m["pak"] = z // จับ key = value || จริงๆแล้วก็คือ m["pak"] = map[int][]string{}

	m["pak"] = make(map[int][]string) //ให้ key m["pak"] = value ก่อน || คือ  make(map[int][]string)
	m["pak"][1] = []string{"z", "x"}
	m["pak"][2] = []string{"zz", "xx"}

	fmt.Println("Add key[pak] in Map m =\n", "key yong =", m["yong"], "\n", "key pak =", m["pak"])

	//============= add m["yong"] ============== เพิ่มข้อมูลลงใน m["yong"]
	m["yong"][3] = []string{"n", "nn"}
	m["yong"][4] = []string{"g", "gg"}

	fmt.Println("Add value in key[yong] = \n", "key yong =", m["yong"], "\n", "key pak =", m["pak"])
	//============= add m["pak"][] ===========
	m["pak"][55][0] = "p" + "ak"
	value := m["pak"][55]
	value = append(value, "111")
	fmt.Println("value in[pak][55]", value)

	m["pak"][55] = append(m["pak"][55], "222")
	fmt.Println("value in pak", m["pak"])

	// Println(value , m["pak"][55] ) ค่าไม่เท่ากัน เพราะจริงๆแล้วคือคนละตัว

	fmt.Println("Add value in key[pak] :in arr[] = \n", "key yong =", m["yong"], "\n", "key pak =", m["pak"])

	_ = m

}
