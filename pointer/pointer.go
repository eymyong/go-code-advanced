package pointer

import "fmt"

func Pointer1() {

	i := 69
	addOnePtr(&i)              //ส่ง i เข้ามา
	fmt.Println(i)             // 70
	fmt.Println(addOnePtr(&i)) //return ถ้าfuncไม่ประกาศreturn จะPrintln ไม่ได้

}

func changeValue(n *int) { // สร้าง'func' ที่รับ 'n *int' แล้วเปลี่ยนค่าโดยที่ไม่เปลี่ยนที่อยู่
	*n = 70
}

func addOnePtr(i *int) int {
	*i += 1
	return *i
}

//============================================================
func Pointer2() {
	s := 8
	t := 9

	pS := &s       // ประกาศ `pS` มารับ address ของ s ซึ่ง `pS` เป็น type *int
	pT := &t       // ประกาศ `pT` มารับ address ของ  t ซึ่ง `pT` เป็น type *int
	addPtr(pS, pT) // return 17

	pPT := &pT          // ประกาศ `pPT` มารับ address ของ pT ซึ่ง `pPT` เป็น type **int
	addPtrHard(pS, pPT) // return 17
}

func addPtr(a, b *int) int {
	vA := *a // vA เอา value ของ a ออกมา ซึ่ง `vA` เป็น type int
	vB := *b //  vB deref `b` => `vB` type int ---- type of `vB` is int
	return vA + vB
}

func addPtrHard(a *int, b **int) int {
	vA := *a
	// `vB` เอา value ของ b ออกมา ซึ่ง `vB` เป็น type *int || ที่เป็น *int ไม่ใช่ int เพราะ `b`ที่รับเข้ามาเป็น `type **int` => deref ออกมาจึงเป็น *int
	vB := *b
	// `vVB` deref `vB` ซึ่ง `vVB`เป็นtype int เนื่องจาก(`vB`เป็นtype *int เพราะถูก deref มาแล้วครึ่งนึ่งจากด้าบน)
	vVB := *vB
	// จากที่ผ่านมาทั้งหมดจึงทำให้ a,b เป็น type int ทั้งคู่แล้ว จึงสามารถน้ำข้อมูลมาบวกกันได้
	return vA + vVB
}

//====================================================================

func Pointer3() {

	p1 := player{name: "yong", level: 1}

	fmt.Println("before level =", p1.level)
	fmt.Println("before address p1 =", &p1.level)
	fmt.Println("===")

	p2 := p1.upLevel1()
	fmt.Println(p2)

	fmt.Println("after level =", p1.level)
	fmt.Println("after address p1 =", &p1.level)

}

type player struct {
	name  string
	level int
}

func (p *player) upLevel1() player { // รับ `*player` || p.level += 1 || return person ที่ value ถูกเปลี่ยนจาก address ไม่ใช่การ Copy
	p.level += 1
	return *p
}

func (p *player) upLevel2() *player { // รับ `*player` || p.level += 1 || return `*person`` ซึ่ง value จะออกไปเป็น &(value)
	p.level += 1 // change p in main
	pCopy := *p  // copy from p in main to p2
	return &pCopy

}
func (p player) upLevel3() player { // เปลี่ยน value in func || แต่ไม่ได้เปลี่ยนที่ address ทำให้พอออกจาก func player.level ยังไม่ถูกเปลี่ยน

	p.level += 1 // copy
	return p     // copy
}

//=============================================================================
//     การบ้าน    //
type Kuy struct{}

func eym(hee **[]int) {
	w := []int{1, 2, 3, 4}
	h := 69
	//t := &h
	w = append(w, h)
	//fmt.Println(w)
	//ww := &w

	//=== in main ===
	// var v **[]int

	// eym(v)
	// fmt.Println() // [1,2,3,4,69]

}
func kuy() {

	//=== in main ===
	// var h []string
	// //kuy(h, 2)
	// fmt.Println(h) //ใส่int 2 ต้องมี [hee,hee] , ใส่int 3 ต้องมี [hee,hee,hee] ,ถ้าintติด- [kuy]

}

//============================================

func Pointer4() {

	s := "mond"
	t := "art"

	ss := &s
	tt := &t

	swapHard(&ss, &tt)

}

func swapHard(a **string, b **string) { // สลับ value โดยที่ ที่อยู่เดิม || วิธีคิดคือเราจะสลับvalue ก็ต้องเข้าถึงvalueให้ได้ก่อน
	fmt.Println("a ก่อน", a)        //&a
	fmt.Println("a value ก่อน", *a) // value a
	fmt.Println("b ก่อน", b)        //&a
	fmt.Println("b value ก่อน", *b)
	fmt.Println("====")
	//newA := a
	newValueA := *b   // สร้างตัวแปรใหม่ = value ของ b ซึ่ง newValueA เป็น type *int
	nnA := &newValueA // สร้างตัวแปรใหม่เป็น *newValueA ซึ่ง nnA เป็น type **int
	newValueB := *a
	nnB := &newValueB

	*a = *nnA // value ของ a = value ของ *nnA // deref`*nnA` = *b
	*b = *nnB

	fmt.Println("a หลัง", a)
	fmt.Println("a value หลัง", *a)
	fmt.Println("b หลัง", b)
	fmt.Println("b value หลัง", *b)
}

//=============================================================//
