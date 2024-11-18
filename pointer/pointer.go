package main

import "fmt"

func change(original *int, value int) {
	*original = value

}
func main() {
	var numberA int = 10
	//var numberB *int = &numberA

	fmt.Println("nilai variabel   :", numberA)
	fmt.Println("alamat pointer :", &numberA)

	var number = 15
	fmt.Println("nilai awal :", number)

	change(&number, 30)
	fmt.Println("nilai setelah perubahan  :", number)
}
