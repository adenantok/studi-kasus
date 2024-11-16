package main

import "fmt"

func main() {
	//if else
	fmt.Println("if else")
	nilai := 70

	if nilai >= 85 {
		fmt.Println("nilai: ", nilai, "\nklarifikasi: baik")
	} else if nilai >= 70 {
		fmt.Println("nilai: ", nilai, "\nklarifikasi: cukup")
	} else if nilai >= 50 {
		fmt.Println("nilai: ", nilai, "\nklarifikasi: kurang")
	} else if nilai < 50 {
		fmt.Println("nilai: ", nilai, "\nklarifikasi: jelek")
	}
	fmt.Println("\nswitch case")
	//switch case
	input := 1

	switch input {
	case 1:
		fmt.Println("input: ", input, "\nhari: senin ")
	case 2:
		fmt.Println("input: ", input, "\nhari: selasa ")
	case 3:
		fmt.Println("input: ", input, "\nhari: rabu ")
	case 4:
		fmt.Println("input: ", input, "\nhari: kamis ")
	case 5:
		fmt.Println("input: ", input, "\nhari: jumat ")
	case 6:
		fmt.Println("input: ", input, "\nhari: sabtu ")
	case 7:
		fmt.Println("input: ", input, "\nhari: minggu ")
	default:
		fmt.Println("tidak ada hari ke:", input)
	}
}
