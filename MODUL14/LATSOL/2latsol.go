package main

import "fmt"

func main() {
	var input int
	var prima bool = true

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&input)

	if input <= 1 {
		fmt.Print("Bukan Bilangan Prima")
		return
	}

	for i := 2; i < input; i++ {
		if input%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
