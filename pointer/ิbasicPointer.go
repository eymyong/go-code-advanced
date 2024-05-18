package pointer

import "fmt"

// pointer  จะมี [ '*'(ตัวชี้) ] , [ '&'(ที่อยู่) ]
// '*' เรียกว่า star , '&' เรียกว่าแอ่น

//1)  `*` หรือ  star ที่อยู่หน้า ตัวแปร  >>  '*input'  จะเป็นการ deref(คือการเรียกหาค่าที่อยู่ในตัวแปรออกมา)
//2)  `*` หรือ  star ที่อยู่หน้า type   >>  '*string' จะเป็นการ ref(คือการชี้ไปหาที่อยู่ของtypeที่ถูกชี้)

//
//ตัวแปร `*`(คือ pointer ที่เป็นตัวชี้ หรือเรียกว่า star) || ไม่ได้เอาไว้เก็บค่า string,int,float64 || แต่เอาไว้ชี้ไปที่ตำแหน่งที่อยู่ คือ address
//=== Ex1 ====
// var a int = 10
// b := &a	 // สร้าง`b`เป็น(type *int) คือ `b` ชี้ไปหาที่อยู่ของ int
// var c **int	 // สร้าง`c`เป็น(type **int) คือ `c` ชี้ไปหาที่อยู่ของ *int ซึ่งก็คือ `b`
// var d ***int	 // สร้าง`d`เป็น(type ***int) คือ `d` ชี้ไปหาที่อยู่ของ **int ซึ่งก็คือ `c`

// `b` จะไม่สามารถ = `a`ได้ || เพราะ `a` เป็น (type int) แต่ `b` เป็น (type *int)
//ถ้าจะให้ `b` เข้าถึง value ของ `a` ได้นั้น || จะต้องทำให้ `b` ซึ่งเป็น (type*int) ชี้ไปหา `a`(type int) || โดยการใส่`&`ข้างหน้า`a`(b = &a) เพื่อที่จะบอกว่า`*` ชี้มาตรง `&`
//และจะ deref ได้ด้วยการใส่ `*` หน้าตัวแปร(*b)
// b = &a  	//  ถ้าปริ้น `b`หมายความว่า  value `b` = 0xc00000a0a8
// c = *b			// `*b` *อยู่หน้าตัวแปร จะ deref value = 10 || ตามข้อ 1) line 8

//ถ้าจะให้ `c` เข้าถึง value ของ `a`ซึ่ง value `a` = 10 || จะต้องทำให้ `c` ซึ่งเป็น (type**int) ชี้ไปหา `b`(type *int) || คือการใส่ `&` ซึ่งก็คือ(c =&b)
//??ทำไมถึงชี้ไปที่ &b ไม่ชี้ไปหา &a || เพราะว่า`c`**int =>*int(ซึ่งก็คือ`b`) || และมี(b =&a อยู่แล้ว) || (c = &&a) || หมายความว่า `c` เป็น *b >> ซึ่ง *b = &a >> ซึ่ง &a = 10 ||
//และจะ deref ได้ด้วยการใส่ `**` หน้าตัวแปร(**c)
// 	c = &b 		// ถ้าปริ้น `c`หมายความว่า  value `c` = 0xc000058028
// d = **c		// `**c` *อยู่หน้าตัวแปร จะ deref value = 10 || ตามข้อ 1) line 8

func BasicPointer() {
	var a int
	a = 10
	var b *int
	b = &a
	c := &b

	fmt.Println("a =", a)
	fmt.Println("*b =", *b)
	fmt.Println("**c =", **c)
	fmt.Println("===")

	fmt.Println("b =", b)
	fmt.Println("&a =", &a)
	fmt.Println("*c", *c)
	fmt.Println("===")

	fmt.Println("c =", c)
	fmt.Println("&b =", &b)
	fmt.Println("===")
}

// func ex(a *int) คือfuncรับ 'a type *int' ==> เรียกfunc จะต้องเป็น ex(&a)
// s := 10, a := &s, b = *a, b = 10,  ==>  ใส่'*ข้างหน้าตัวแปรเพื่อจะ deref &ที่อยู่หน้าตัวแปร
// 101      102     103
//เช็คย้อนกลับว่าตัวแปรที่เราใช้อยู่เป็น typeอะไร ==> นับย้อน&, e

//

func ChangeInt(i ***int) *int {
	a := **i
	return a
}

// j := 10      // 1
// p1 := &j     // 2           // p1 = 0x7
// p2 := foo(j) // 1           // p2 = 0x8

// func foo(j int){

// }

// fmt.Printf("p1 %p p2 %p\n", p1, p2)
// *p2 = 7
// fmt.Println("j", j)

// var p3 *int // nil deref = panic
// var p4 *int = new(int) // จะไปหาที่ว่าง แล้วสร้าง default int (0) แล้วเอาที่อยู่อันนั้นมาให้ ซึ่งถ้า deref `p4` = 0
// fmt.Println(p3, p4, *p3)
