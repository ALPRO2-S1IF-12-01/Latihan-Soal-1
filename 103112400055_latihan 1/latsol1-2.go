package main

import "fmt"

func main() {
	fmt.Println("\nContoh for sebagai while:")
	counter := 1
	for counter <= 5 {
		fmt.Printf("Iterasi ke-%d\n", counter) // (4) Lengkapi bagian ini untuk mencetak "Iterasi ke-X"
		counter++
	}
}
