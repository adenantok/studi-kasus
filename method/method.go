package main

import "fmt"

type Orang struct {
	nama string
	umur int
}

func (o Orang) Tampilnama() {
	fmt.Println("nama saya", o.nama)

}

func (u *Orang) Tampilumur(umur int) {
	u.umur = umur
}

func main() {
	var orang = Orang{nama: "alice", umur: 12}
	orang.Tampilnama()

	orang.Tampilumur(33)
	fmt.Println("umur", orang.umur)
}
