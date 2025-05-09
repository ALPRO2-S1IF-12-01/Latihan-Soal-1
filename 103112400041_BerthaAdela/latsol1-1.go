//BERTHA ADELA
//103112400041
package main

import (
	"fmt"
)

func main() {
	// 1. Struktur kondisional if-else
	nilai := 85
	fmt.Println("Contoh if-else:")
	if nilai >= 90 {
		fmt.Println("Nilai A") // (1) Lengkapi bagian ini untuk mencetak "Nilai A"
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("Nilai C") // (2) Lengkapi bagian ini untuk mencetak "Nilai C"
	} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E") // (3) Lengkapi bagian ini untuk mencetak "Nilai E"
	}
	// 2.
	counter := 1
	for counter <= 5 {
		fmt.Print(counter) // (4) Lengkapi bagian ini untuk mencetak "Iterasi ke-X"
		counter++
	}

	// 3. Struktur perulangan for dengan range
	fmt.Println("\nContoh for dengan range:")
	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for _______, item := range buah { // (5) Lengkapi bagian ini agar mencetak indeks dan nama buah
	_________________________ // (6) Lengkapi bagian ini untuk mencetak "Buah pada index X adalah Y"
	}
	// 4. Struktur switch-case

	fmt.Println("\nContoh switch-case:")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari kerja") // (7) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Selasa":
		fmt.Println("Hari kerja")
	case "Rabu":
		fmt.Println("Hari kerja lengkap") // (8) Lengkapi bagian ini agar hari kerja lengkap
	case "Kamis":
		fmt.Println("Hari kerja")
	case "Jumat":
		fmt.Println("Hari kerja") // (9) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur") // (10) Lengkapi bagian ini untuk mencetak "Hari libur"
	default:
		fmt.Println("Hari tidak valid")
	}


}