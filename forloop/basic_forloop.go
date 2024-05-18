package forloop

import "fmt"

func ForLoop() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// for i := 1; i <= 10; i-- {
	// 	fmt.Printf("Number: %d\n", i)
	// }

}

func DoWhileLoop() {
	i := 1
	for {
		fmt.Printf("Number: %d\n", i)
		i++
		if i > 5 {
			break
		}
	}
}

func WhileLoop() {
	i := 1
	for i < 10 {
		fmt.Printf("Number: %d\n", i)
		i++
	}
}
