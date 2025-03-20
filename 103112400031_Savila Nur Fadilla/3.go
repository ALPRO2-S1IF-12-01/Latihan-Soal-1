package main

import (
	"fmt"
)

// Prosedur sederhana tanpa parameter
func tampilkanHeader() {
	fmt.Println("=========================")
	fmt.Println("        PROGRAM MAHASISWA        ")
	fmt.Println("=========================")
}

// Prosedur dengan parameter value
func tampilkanInfo(nama string, nim string, jurusan string) {
	fmt.Println("Informasi Mahasiswa:")
	fmt.Printf("Nama    : %s\n", nama)
	fmt.Printf("NIM    : %s\n", nim)
	fmt.Printf("Jurusan : %s\n", jurusan)
}

// Prosedur dengan parameter pointer
func ubahNilai(nilai *int) {
	*nilai += 10
	fmt.Printf("Nilai setelah diubah: %v\n", *nilai)
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
		fmt.Println(matkul, nilai)
	}
}

// Prosedur dengan slice parameter
func tampilkanDaftarMahasiswa(daftar []string) {
	fmt.Println("Daftar Mahasiswa:")
	for i, nama := range daftar {
		fmt.Println(i, nama)
	}
}

func main() {
	// Memanggil prosedur tanpa parameter
	tampilkanHeader()
	
	// Memanggil prosedur dengan parameter value
	tampilkanInfo("Ani", "87654321", "Informatika")
	
	// Memanggil prosedur dengan parameter pointer
	nilai := 75
	fmt.Printf("Nilai awal: %d\n", nilai)
	ubahNilai(&nilai)
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
}
