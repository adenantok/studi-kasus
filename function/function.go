package main

import "fmt"

func bilangan(bilangan1, bilangan2 int) {
	jumlah := bilangan1 + bilangan2
	fmt.Println("jumlah", jumlah)
}

func kuadrat(param int) int {
	kuadrat := param * param
	return kuadrat
}

func main() {
	bilangan(2, 2)

	bil := 4
	kuadrat := kuadrat(bil)
	fmt.Println("bilangan kuadrat dari", bil, ":", kuadrat)

	//balok
}
