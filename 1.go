package main

import (
	"fmt"
)

func main() {
	//1. struktur if-else
	nilai := 85
	fmt.Println("Contoh if-else:")
	if nilai >= 90 {
		fmt.Print("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("Nilai C")
	} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		fmt.Print("Nilai E")
	}

	//2. for loop
	fmt.Println("\nContoh for sebagai while:")
	counter := 1
	for counter <= 5 {
		fmt.Printf("Iterasi ke-%d\n", counter)
		counter++
	}

	//3. for loop with range
	fmt.Println("\nContoh for dengan range:\n")
	buah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}
	for X, Y := range buah {
		fmt.Printf("Buah pada index %d adalah %s\n",X, Y )
	}

	//4. switch case
	fmt.Println("\nContoh switch-case:")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Print("Hari Kerja")
	case "Selasa":
		fmt.Println("Hari kerja")
	case "Rabu":
		fmt.Print("Hari Kerja") 
	case "Kamis":
		fmt.Println("Hari kerja")
	case "Jumat":
		fmt.Print("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Print("Hari Libur")
	default:
		fmt.Println("Hari tidak valid")
	}
}


