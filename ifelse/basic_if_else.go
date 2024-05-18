package ifelse

import "fmt"

//condition statement

func TestIfelse() {
	number := 75
	grade(number)
	foo()

}

func foo() {
	num1 := 5
	num2 := 10
	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println(sumNum)
	}

}

func Buy(score int) {
	if score > 50 {
		fmt.Println("ซื้อได้")
	} else {
		fmt.Println("ซื้อไม่ได้")
	}
}

func grade(score int) {
	if score >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else if score >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if score >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if score >= 50 {
		fmt.Println("คุณได้เกรด D")
	} else {
		fmt.Println("คุณตก")
	}

}

func grade2(score int) { // ข้อควรระวัง ต้องเรียงลำดับ condition ให้ถูกต้อง เพราะถ้าเข้าเงื่อนไขไปแล้ว จะไม่ลงมาทำต่อ
	if score >= 50 { //ต่อ) สมมุติถ้า score = 85 ตามจริงต้องได้ grade `A` แต่กลับได้แค่ grade `D`
		fmt.Println("คุณได้เกรด D")
	} else if score >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if score >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if score >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else {
		fmt.Println("คุณตก")
	}

}

func grade3() {
	var score int
	fmt.Println("คะแนน")
	fmt.Scanf("%d", &score) // "%d", &score คือทำให้พอ run แล้ว สามารถพิมพ์ input เข้าไปใน terminal ได้และ func จะทำงานต่อ
	fmt.Println("score = ", score)
	if score >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else if score >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if score >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if score >= 50 {
		fmt.Println("คุณได้เกรด D")
	} else {
		fmt.Println("คุณตก")
	}

}
