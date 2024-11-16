package main

import "fmt"

func main() {
	var bilBul [5]int
	bilBul[0] = 1
	bilBul[1] = 2
	bilBul[2] = 3
	bilBul[3] = 4
	bilBul[4] = 5

	fmt.Println(bilBul[0], bilBul[1], bilBul[2], bilBul[3], bilBul[4])

	//Loop array with for
	fmt.Println("\nLoop array with for")
	for i := 1; i <= 5; i++ {
		fmt.Println("elemen: ", i)
	}

	//Loop array with range
	fmt.Println("\nLoop array with range")

	for i, v := range bilBul {
		fmt.Printf("Index %d :%d\n", i, v)
	}

	//Menggunakan variabel  _ di for range
	fmt.Println("\nLMenggunakan variabel  _ di for range")

	for _, v := range bilBul {
		fmt.Printf("%d\n", v)
	}
}
