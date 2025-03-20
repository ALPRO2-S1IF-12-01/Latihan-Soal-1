package main

import "fmt"

func main() {

	fmt.Println("\nContoh for dengan range:")

	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for _, item := range buah { // (5) Lengkapi bagian ini agar mencetak indeks dan nama buah
		fmt.Println(buah, item) // (6) Lengkapi bagian ini untuk mencetak "Buah pada index X adalah Y"
	}

}
