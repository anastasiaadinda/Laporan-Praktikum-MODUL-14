package main

import "fmt"

func main() {
	var input int
	var ulang int

	fmt.Print("Masukan bilangan positif: ")
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {
		if i%2 != 0 {
			ulang++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", ulang)
}
