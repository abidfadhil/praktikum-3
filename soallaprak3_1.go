package main

import (
	"fmt"
)

func main() {
	var totalBelanja, diskon int

	fmt.Println("Masukkan total belanja awal:")
	fmt.Scan(&totalBelanja)

	fmt.Println("Masukkan diskon (dalam persen):")
	fmt.Scan(&diskon)

	totalSetelahDiskon := totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println("Total belanja setelah diskon:", totalSetelahDiskon)
}
