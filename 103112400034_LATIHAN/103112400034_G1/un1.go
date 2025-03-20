package main

import (
    "fmt"
)

func main() {
    var n int 
	fmt.Println("n: ")
	fmt.Scan(&n)
    nilai := 85

   
    if nilai >= 90 {
        fmt.Println("Nilai A")
    } else if nilai >= 80 {
        fmt.Println("Nilai B")
    } else if nilai >= 70 {
        fmt.Println("Nilai C")
    } else if nilai >= 60 {
        fmt.Println("Nilai D")
    } else {
        fmt.Println("Nilai E")
    }

    // 2. Struktur perulangan for (seperti while)
    fmt.Println("\nContoh for sebagai while:")
    counter := 1
    for counter <= 5 {
        fmt.Printf("Iterasi ke-%d\n", counter)
        counter++
    }
}