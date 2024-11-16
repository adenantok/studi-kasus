package main

import "fmt"

func main() {
	//Initialize Slice
	var buah = []string{"apel", "pisang", "ceri"}

	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}

	//Initialize slice from array
	fmt.Println("\nInitialize slice from array")
	var angka = []int{1, 2, 3, 4, 5}
	var newAngka = angka[1:4]
	for i := 0; i < len(newAngka); i++ {
		fmt.Println(newAngka[i])
	}

	fmt.Println("Panjang: ", len(newAngka))

	// Append
	fmt.Println("\nAppend")
	var appendBuah = append(buah, "durian")
	for i := 0; i < len(appendBuah); i++ {
		fmt.Println(appendBuah[i])
	}

	//initialize map
	fmt.Println("\ninitialize map")
	initMap := map[string]int{
		"alice":   25,
		"bob":     30,
		"charlie": 35,
	}

	for key, val := range initMap {
		fmt.Println(key, " ", val)
	}

	//Delete map item
	fmt.Println("\nDelete map item")
	delete(initMap, "charlie")

	for key, val := range initMap {
		fmt.Printf("%s: %d\n", key, val)
	}

	_, exist := initMap["charlie"]
	if exist {
		fmt.Printf("kunci ada: ada")
	} else {
		fmt.Println("kunci ada: tidak")
	}
}
