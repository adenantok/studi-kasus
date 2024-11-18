package main

import "fmt"

type Orang struct {
	nama string
	usia int8
}

type Siswa struct {
	lokasi string
	Orang
}

func main() {
	var orang Orang

	orang.nama = "alice"
	orang.usia = 20

	fmt.Println("nama: ", orang.nama)
	fmt.Println("usia: ", orang.usia)

	//object struct
	fmt.Println("\nobject struct")
	var orang1 = Orang{
		nama: "bob",
		usia: 30,
	}

	fmt.Println("nama: ", orang1.nama)
	fmt.Println("usia: ", orang1.usia)

	//Embed a struct
	fmt.Println("\nEmbed a struct")
	var orang2 = Siswa{}
	orang2.nama = "charli"
	orang2.usia = 33
	orang2.lokasi = "jakarta"

	fmt.Println("nama: ", orang2.nama)
	fmt.Println("usia: ", orang2.usia)
	fmt.Println("lokasi: ", orang2.lokasi)

	//Slice of struct
	fmt.Println("\nSlice of struct")
	var orang3 = []Orang{
		{nama: "dedi"},
		{nama: "rizal"},
		{nama: "indah"},
	}

	for _, tampilOrang3 := range orang3 {
		fmt.Println("nama: ", tampilOrang3.nama)
	}
}
