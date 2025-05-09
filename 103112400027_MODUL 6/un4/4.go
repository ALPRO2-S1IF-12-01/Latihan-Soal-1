package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks (akhiri dengan titik): ")
	i := 0
	for i < NMAX {
		ch, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Terjadi kesalahan saat membaca input:", err)
			break
		}
		if ch == '.' {
			break
		}
		(*t)[i] = ch
		i++
	}
	*n = i
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks Asli       : ")
	cetakArray(tab, m)

	if palindrome(tab, m) {
		fmt.Println("Palindrome?     : true")
	} else {
		fmt.Println("Palindrome?     : false")
	}

	balikanArray(&tab, m)
	fmt.Print("Teks Terbalik   : ")
	cetakArray(tab, m)
}
