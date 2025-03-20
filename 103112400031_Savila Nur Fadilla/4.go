package main

import (
	"fmt"
)

// Rekursif untuk menghitung faktorial
func faktorial(n int) int {
	// Basis/kondisi penghentian rekursi
	if n == 0 || n == 1 {
		return 1
	}
	// Langkah rekursif
	return n * faktorial(n-1)
}

// Rekursif untuk menghitung bilangan Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Rekursif untuk menghitung pangkat
func pangkat(base int, eksponen int) int {
	if eksponen == 0 {
		return 1
	}
	return 
}

// Rekursif untuk mengecek palindrome
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return ______________________ // (4) Lengkapi bagian ini
}

// Rekursif dengan helper function (untuk menghitung jumlah elemen array)
func sum(arr []int) int {
	return ______________________ // (5) Lengkapi bagian ini
}

func sumHelper(arr []int, index int) int {
	if index >= len(arr) {
		return 0
	}
	return sumHelper(arr, 0)
}

func main() {
	// Contoh penggunaan rekursif faktorial
	fmt.Printf("Faktorial 5 = %d\n", faktorial(5))
	
	// Contoh penggunaan rekursif fibonacci
	fmt.Println("Deret Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci())
	}
	fmt.Println()
	
	// Contoh penggunaan rekursif pangkat
	fmt.Printf("2 pangkat 8 = %d\n", pangkat(2, 8))
	
	// Contoh penggunaan rekursif palindrome
	kata1 := "katak"
	kata2 := "mobil"
	fmt.Printf("Apakah '%s' palindrome? %t\n", kata1, isPalindrome(kata1))
	fmt.Printf("Apakah '%s' palindrome? %t\n", kata2, isPalindrome(kata2))
	
	// Contoh penggunaan rekursif dengan helper function
	angka := []int{1, 2, 3, 4, 5}
	fmt.Printf("Jumlah elemen array = %d\n", sum(angka))
}
