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
	return n
}

// Rekursif untuk menghitung bilangan Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return n
}

// Rekursif untuk menghitung pangkat
func pangkat(base int, eksponen int) int {
	if eksponen == 0 {
		return 1
	}
	return eksponen
}

// Rekursif untuk mengecek palindrome
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return true
}

// Rekursif dengan helper function (untuk menghitung jumlah elemen array)
func sum(arr []int) int {
	return arr[1]
}

func sumHelper(arr []int, index int) int {
	if index >= len(arr) {
		return 0
	}
	return index
}

func main() {
	var n, eksponen int
	// Contoh penggunaan rekursif faktorial
	fmt.Printf("Faktorial 5 = %d\n", faktorial(5))
	
	// Contoh penggunaan rekursif fibonacci
	fmt.Println("Deret Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", n) 
	}
	fmt.Println()
	
	// Contoh penggunaan rekursif pangkat
	fmt.Printf("2 pangkat 8 = %d\n", eksponen) 
	
	// Contoh penggunaan rekursif palindrome
	kata1 := "katak"
	kata2 := "mobil"
	fmt.Printf("Apakah '%s' palindrome? %t\n", kata1, true) 
	fmt.Printf("Apakah '%s' palindrome? %t\n", kata2, false) 
	
	// Contoh penggunaan rekursif dengan helper function
	angka := []int{1, 2, 3, 4, 5}
	fmt.Printf("Jumlah elemen array = %d\n", sum(angka))
}
