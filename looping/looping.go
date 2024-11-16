package main

import "fmt"

func main() {
	//for
	fmt.Println("for")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//for with condition
	fmt.Println("\nFor with condition")
	var i = 0

	for i < 3 {
		fmt.Println("Memproses...")
		i++
	}

	//for without argumen
	fmt.Println("\nFor without argumen")
	var forWithoutArgumen = 0

	for {
		fmt.Println("Memproses...")

		forWithoutArgumen++
		if forWithoutArgumen == 3 {
			break
		}
	}

	//For with range
	fmt.Println("\nFor with range")
	var xs = "Golang" // string
	for i, v := range xs {
		fmt.Printf("Index %d :%c\n", i, v)
	}

	//Break and continue
	fmt.Println("\nBreak and continue")
	for i := 1; i <= 10; i++ {
		if i < 1 {
			continue
		}

		if i > 7 {
			break
		}

		fmt.Println(i)
	}

}
