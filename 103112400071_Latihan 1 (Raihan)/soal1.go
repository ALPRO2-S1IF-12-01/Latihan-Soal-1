package main

import (
	"fmt"
)

func main() {
	// 1. Struktur kondisional if-else
	nilai := 85
	fmt.Println("Contoh if-else:")
	if nilai >= 90 {
		_________________________ // (1) Lengkapi bagian ini untuk mencetak "Nilai A"
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		_________________________ // (2) Lengkapi bagian ini untuk mencetak "Nilai C"
	} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		_________________________ // (3) Lengkapi bagian ini untuk mencetak "Nilai E"
	}

	// 2. Struktur perulangan for (seperti while)
	fmt.Println("\nContoh for sebagai while:")
	counter := 1
	for counter <= 5 {
		_________________________ // (4) Lengkapi bagian ini untuk mencetak "Iterasi ke-X"
		counter++
	}

	// 3. Struktur perulangan for dengan range
	fmt.Println("\nContoh for dengan range:")
	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for ________, item := range buah { // (5) Lengkapi bagian ini agar mencetak indeks dan nama buah
		_________________________ // (6) Lengkapi bagian ini untuk mencetak "Buah pada index X adalah Y"
	}

	// 4. Struktur switch-case
	fmt.Println("\nContoh switch-case:")
	hari := "Senin"
	switch hari {
	case "Senin":
		_________________________ // (7) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Selasa":
		fmt.Println("Hari kerja")
	case "Rabu":
		_________________________ // (8) Lengkapi bagian ini agar hari kerja lengkap
	case "Kamis":
		fmt.Println("Hari kerja")
	case "Jumat":
		_________________________ // (9) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Sabtu", "Minggu":
		_________________________ // (10) Lengkapi bagian ini untuk mencetak "Hari libur"
	default:
		fmt.Println("Hari tidak valid")
	}
}
