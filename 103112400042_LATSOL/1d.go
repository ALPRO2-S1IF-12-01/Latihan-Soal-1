package main

import "fmt"

func main() {
	fmt.Println("\nContoh switch-case:")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("HARI KERJA") // (7) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Selasa":
		fmt.Println("Hari kerja")
	case "Rabu":
		fmt.Println("hari kerja") // (8) Lengkapi bagian ini agar hari kerja lengkap
	case "Kamis":
		fmt.Println("Hari kerja")
	case "Jumat":
		fmt.Println("HARI KERJA") // (9) Lengkapi bagian ini untuk mencetak "Hari kerja"
	case "Sabtu", "Minggu":
		fmt.Println("hari libur") // (10) Lengkapi bagian ini untuk mencetak "Hari libur"
	default:
		fmt.Println("Hari tidak valid")
	}
}
