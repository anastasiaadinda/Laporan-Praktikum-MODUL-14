package main

import "fmt"

func main() {
	//urutan warna benar
	warnaBenar := [4]string{"Merah", "Kuning", "Hijau", "Ungu"}

	var warna string
	hasil := true

	//perulangan untuk 5 percobaan
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		for j := 0; j < 4; j++ {
			fmt.Scan(&warna)
			if warna != warnaBenar[j] {
				hasil = false
			}
		}
	}

	if hasil {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
