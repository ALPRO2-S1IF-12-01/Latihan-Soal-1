//103112400032_Daffa Tsaqifna Fauztsany

package main

import (
	"fmt"
)

// Prosedur sederhana tanpa parameter
func tampilkanHeader() {
	fmt.Println("================================") // (1) Lengkapi untuk mencetak garis atas
	fmt.Println("        PROGRAM MAHASISWA        ")
	fmt.Println("================================") // (2) Lengkapi untuk mencetak garis bawah
}

// Prosedur dengan parameter value
func tampilkanInfo(nama string, nim string, jurusan string) {
	fmt.Println("Informasi Mahasiswa:")
	fmt.Printf("Nama    : %s\n", nama)
	fmt.Printf("NIM     : %s\n", nim) // (3) Lengkapi agar mencetak NIM dengan format yang benar
	fmt.Printf("Jurusan : %s\n", jurusan)
}

// Prosedur dengan parameter pointer
func ubahNilai(nilai *int) {
	*nilai += 10
	return nilai // (4) Lengkapi agar mencetak nilai setelah diubah
}

// Prosedur dengan struct parameter
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	Nilai   map[string]int
}

func tampilkanNilai(mhs Mahasiswa) {
	fmt.Printf("Nilai mahasiswa %s:\n", mhs.Nama)
	for matkul, nilai := range mhs.Nilai {
		_________________________ // (5) Lengkapi agar mencetak nama mata kuliah dan nilai
	}
}

// Prosedur dengan slice parameter
func tampilkanDaftarMahasiswa(daftar []string) {
	fmt.Println("Daftar Mahasiswa:")
	for i, nama := range daftar {
		_________________________ // (6) Lengkapi agar mencetak nomor dan nama mahasiswa
	}
}

func main() {
	// Memanggil prosedur tanpa parameter
	fmt.Print(tampilkanHeader()) // (7) Lengkapi agar memanggil prosedur tampilkanHeader

	// Memanggil prosedur dengan parameter value
	fmt.Print(tampilkanInfo(Mahasiswa.Nama, Mahasiswa.NIM, Mahasiswa.Jurusan)) // (8) Lengkapi agar memanggil prosedur tampilkanInfo dengan data yang sesuai

	// Memanggil prosedur dengan parameter pointer
	nilai := 75
	fmt.Printf("Nilai awal: %d\n", nilai)
	nilai = ubahNilai(&nilai) // (9) Lengkapi agar memanggil prosedur ubahNilai dengan parameter yang benar
	fmt.Printf("Nilai akhir: %d\n", nilai)

	// Memanggil prosedur dengan struct parameter
	mhs := Mahasiswa{
		Nama:    "Ani Wijaya",
		NIM:     "87654321",
		Jurusan: "Sistem Informasi",
		Nilai: map[string]int{
			"Algoritma":       85,
			"Basis Data":      90,
			"Pemrograman Web": 78,
			"Struktur Data":   82,
		},
	}
	_________________________ // (10) Lengkapi agar memanggil prosedur tampilkanNilai dengan parameter yang sesuai
}
