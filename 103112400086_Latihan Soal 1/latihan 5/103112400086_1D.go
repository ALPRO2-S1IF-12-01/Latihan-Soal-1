package main

import "fmt"

func main() {

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
