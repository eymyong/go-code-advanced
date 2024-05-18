package homework1

import "fmt"

type fruitColors map[string]string

func HomeworkMethot() {
	f := fruitColors{
		"mango":  "yellow",
		"orange": "orange",
		"banana": "yellow",
		"tangmo": "red",
	}
	fmt.Println("Before f =", f)
	f.addFruit("durian", "yellow")
	fmt.Println("After f =", f)
	fmt.Println("=====")

	colorInF := f.checkColor("durian")
	fmt.Println(colorInF)
	fmt.Println("=====")

	numberColor := f.groupNumberFruits("yellow")
	numberColor2 := f.groupNumberFruits("orange")
	numberColor3 := f.groupNumberFruits("red")
	fmt.Println(numberColor)
	fmt.Println(numberColor2)
	fmt.Println(numberColor3)
	fmt.Println("====")

	fmt.Println(f.groupFruits("yellow"))
	fmt.Println(f.groupFruits("red"))

}

// methot add fruitColors
func (f fruitColors) addFruit(fruitName, color string) {
	f[fruitName] = color
}

//methot ที่รับ fruit และ return color
func (f fruitColors) checkColor(fruitName string) string {
	_, ok := f[fruitName]
	if ok == false {
		panic("No Fruit Name")
	}
	//color := f[fruitName]
	return f[fruitName]
}

//methot ที่เช็คว่าสีนั้นๆมีผลไม้จำนวนกี่ชนิด
func (f fruitColors) groupNumberFruits(color string) map[string]int {
	numberFruit := make(map[string]int)
	for _, v := range f {
		if color == v {
			numberFruit[color]++
		}
	}
	return numberFruit
}

// 		"mango":  "yellow",
// 		"orange": "orange",
// 		"banana": "yellow",
// 		"tangmo": "red",

//methot ที่เช็คว่าในสีนั้นๆ มีผลไม้ชนิดไหนบ้าง
func (f fruitColors) groupFruits(color string) map[string][]string {
	group := make(map[string][]string)
	key := color
	value := group[key]
	for i, v := range f {
		if v == color {
			value = append(value, i)
		}
	}
	group[key] = value

	return group
}
