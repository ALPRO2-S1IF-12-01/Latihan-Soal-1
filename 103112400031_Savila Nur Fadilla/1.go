// Savila Nur Fadilla
// 10311240031

package main

import (
	"fmt"
)

func main() {
	// 1. Struktur kondisional if-else
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

	// 2. Struktur perulangan for (seperti while)
	fmt.Println("\nContoh for sebagai while:")
	counter := 1
	for counter <= 5 {
		fmt.Printf("Iterasi ke-%v\n", counter)
		counter++
	}

	// 3. Struktur perulangan for dengan range
	fmt.Println("\nContoh for dengan range:")
	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for indeks, item := range buah {
		fmt.Printf("Buah pada indeks %v adalah %v\n", indeks, item)
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
