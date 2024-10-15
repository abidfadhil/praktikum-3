package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukan Berat Badan (kg): ")
	fmt.Scan(&bmi)
	fmt.Print("Masukan Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("beratBadan: %.f", beratBadan)
}
