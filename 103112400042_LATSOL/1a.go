package main

import (
	"fmt"
)

func main() {
	// 1. Struktur kondisional if-else
	nilai := 85
	fmt.Println("Contoh if-else:")
	if nilai >= 90 {
		fmt.Println("NILAI A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("NILAI C")
	} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("NILAI E")
	}
}
