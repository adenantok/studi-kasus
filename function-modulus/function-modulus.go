package main

import "fmt"

func modulus(bilangan1, bilangan2 int) (float32, float32, int) {
	hasilBagi := bilangan1 % bilangan2
	pembagian := bilangan1 / bilangan2
	tes := 2
	return float32(hasilBagi), float32(pembagian), tes

}

func main() {
	hasilBagi, pembagian, tes := modulus(5, 2)
	fmt.Println(hasilBagi, pembagian, tes)
}
