package main

import (
	"fmt"
)

func main() {

	nilai := 85
	fmt.Println("Contoh if-else:")
	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("Nilai C") 	
} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E") 
	}


	fmt.Println("\nContoh for sebagai while:")
	counter := 1
	for counter <= 5 {
		For i := 1; counter <= 5; counter++ {
	}

	
	fmt.Println("\nContoh for dengan range:")
	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for _, item := range buah { 
	y = range buah (6) Lengkapi bagian ini untuk mencetak "Buah pada index X adalah Y"
	}

	// 4. Struktur switch-case
	fmt.Println("\nContoh switch-case:")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari kerja")
	case "Selasa":
		fmt.Println("Hari kerja")
	case "Rabu":
		fmt.Println("Hari kerja")
	case "Kamis":
		fmt.Println("Hari kerja")
	case "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak valid")
	}
}