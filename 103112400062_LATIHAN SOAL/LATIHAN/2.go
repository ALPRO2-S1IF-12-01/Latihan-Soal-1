package main

import (
	"fmt"
	"math"
)

// Fungsi dengan parameter dan return value
func hitungLuasLingkaran(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) // (1) Rumus luas lingkaran: π * r^2
}

// Fungsi dengan multiple return values
func minMax(angka []int) (int, int) {
	if len(angka) == 0 {
		return 0, 0 // (2) Return nilai 0 jika array kosong
	}

	min := angka[0]
	max := angka[0]

	for _, nilai := range angka {
		if nilai < min {
			min = nilai // (3) Update nilai minimum jika ditemukan angka yang lebih kecil
		}
		if nilai > max {
			max = nilai // (4) Update nilai maksimum jika ditemukan angka yang lebih besar
		}
	}

	return min, max
}

// Fungsi dengan named return values
func hitungStatistik(angka []float64) (min, max, avg float64) {
	if len(angka) == 0 {
		return 0, 0, 0
	}

	min = angka[0]
	max = angka[0]
	var total float64 = 0

	for _, nilai := range angka {
		if nilai < min {
			min = nilai // (5) Update nilai minimum jika ditemukan angka yang lebih kecil
		}
		if nilai > max {
			max = nilai // (6) Update nilai maksimum jika ditemukan angka yang lebih besar
		}
		total += nilai
	}

	avg = total / float64(len(angka)) // (7) Rata-rata: jumlah total dibagi jumlah elemen
	return                            // implisit return untuk named return values
}

// Fungsi dengan variadic parameter
func jumlahkan(angka ...int) int {
	total := 0
	for _, nilai := range angka {
		total += nilai // (8) Proses penjumlahan
	}
	return total
}

func main() {
	// Contoh penggunaan fungsi dengan return value
	radius := 7.0
	luas := hitungLuasLingkaran(radius) // (9) Panggil fungsi hitungLuasLingkaran dengan parameter yang benar
	fmt.Printf("Luas lingkaran dengan jari-jari %.1f adalah %.2f\n", radius, luas)

	// Contoh penggunaan fungsi dengan multiple return values
	data := []int{23, 45, 12, 67, 34, 8}
	minimal, maksimal := minMax(data) // (10) Panggil fungsi minMax dengan parameter yang benar
	fmt.Printf("Nilai minimum: %d, Nilai maksimum: %d\n", minimal, maksimal)
}