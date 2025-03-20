package main

import (
	"fmt"
)

// Fungsi dengan parameter dan return value
func hitungLuasLingkaran(jariJari float64) float64 {
	return 3.14*jariJari*jariJari // (1) Lengkapi rumus luas lingkaran
}

// Fungsi dengan multiple return values
func minMax(angka []int) (int, int) {
	if len(angka) == 0 {
		return 0, 0, 0 // (2) Pastikan return value yang benar jika array kosong
	}

	min := angka[0]
	max := angka[0]

	for _, nilai := range angka {
		if nilai < min {
			_________________________ // (3) Lengkapi agar min selalu mendapat nilai terkecil
		}
		if nilai > max {
			_________________________ // (4) Lengkapi agar max selalu mendapat nilai terbesar
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
			_________________________ // (5) Lengkapi agar min selalu mendapat nilai terkecil
		}
		if nilai > max {
			_________________________ // (6) Lengkapi agar max selalu mendapat nilai terbesar
		}
		total += nilai
	}

	avg = total / _________________________ // (7) Lengkapi perhitungan rata-rata
	return                                  // implisit return untuk named return values
}

// Fungsi dengan variadic parameter
func jumlahkan(angka ...int) int {
	total := 0
	for _, nilai := range angka {
		_________________________ // (8) Lengkapi proses penjumlahan
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
