package main

import (
	"fmt"
)


func tampilkanHeader() {
	fmt.Print("=====================================") // (1) Lengkapi untuk mencetak garis atas
	fmt.Println("        PROGRAM MAHASISWA        ")
	fmt.Print("=====================================")// (2) Lengkapi untuk mencetak garis bawah
}

// Prosedur dengan parameter value
func tampilkanInfo(nama string, nim string, jurusan string) {
	fmt.Println("Informasi Mahasiswa:")
	fmt.Printf("Nama    : %s\n", nama)
	fmt.Printf("Nim  : %s\n", nim)//dengan format yang benar
	fmt.Printf("Jurusan : %s\n", jurusan)
}

// Prosedur dengan parameter pointer
func ubahNilai(nilai *int) {
	*nilai += 10
	 return nilaingkapi agar mencetak nilai setelah diubah
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
		
	}
}

// Prosedur dengan slice parameter
func tampilkanDaftarMahasiswa(daftar []string) {
	fmt.Println("Daftar Mahasiswa:")
	for i, nama := range daftar {
		fmt.Print(nama,nim)// (6) Lengkapi agar mencetak nomor dan nama mahasiswa
	}
}

func main() {
	// Memanggil prosedur tanpa parameter
	tampilkanHeader()
	
	// Memanggil prosedur dengan parameter value
	tampilkanInfo("ani wijaya","87654321","sistem informasi")
	
	// Memanggil prosedur dengan parameter pointer
	nilai := 75
	fmt.Printf("Nilai awal: %d\n", nilai)
	ubahNilai(9) Lengkapi agar memanggil prosedur ubahNilai dengan parameter yang benar
	fmt.Printf("Nilai akhir: %d\n", nilai)
	
	// Memanggil prosedur dengan struct parameter
	mhs := Mahasiswa{
		Nama:    "Ani Wijaya",
		NIM:     "87654321",
		Jurusan: "Sistem Informasi",
		Nilai: map[string]int{
			"Algoritma":          85,
			"Basis Data":         90,
			"Pemrograman Web":    78,
			"Struktur Data":      82,
		},
	}
	tampilkanNilai(mhs)
