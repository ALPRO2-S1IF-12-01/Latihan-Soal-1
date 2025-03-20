package main

// Fungsi dengan parameter dan return value
func hitungLuasLingkaran(jariJari float64) float64 {
	return math.Pi * jariJari * jariJari// (1) Lengkapi rumus luas lingkaran
}

// Fungsi dengan multiple return values
func minMax(angka []int) (int, int) {
	if len(angka) == 0 {
		return 0, 0// (2) Pastikan return value yang benar jika array kosong
	}

	min := angka[0]
	max := angka[0]

	for _, nilai := range angka {
		if nilai < min {
			min = nilai// (3) Lengkapi agar min selalu mendapat nilai terkecil
		}
		if nilai > max {
			max = nilai // (4) Lengkapi agar max selalu mendapat nilai terbesar
		}
	}

	return min, max
}
