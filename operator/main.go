package main

import "fmt"

func main() {
	var panjang int = 10
	var lebar int = 5

	//menghitung luas
	var luas = panjang * lebar

	//menghitung keliling
	var keliling = 2 * (panjang + lebar)

	//memanggil luas
	fmt.Println("Luas persegi panjang:", luas, keliling)

	//memanggil keliling
	fmt.Println("Keliling :", keliling)

	//perbandingan
	var angkaPertama = 15
	var angkaKedua = 20

	if angkaPertama < angkaKedua {
		fmt.Println(angkaPertama, "lebih kecil daripada ", angkaKedua)
	} else if angkaPertama > angkaKedua {
		fmt.Println(angkaPertama, "lebih besar dari ", angkaKedua)
	}

	//logika
	var usia = 19
	//var wni = true
	var wargaNegara = "malay"

	if usia >= 18 && wargaNegara == "indonesia" {
		fmt.Println("memenuhi syarat memilih : Ya")
	} else {
		fmt.Println("memenuhi syarat memilih : tidak")
	}
}
