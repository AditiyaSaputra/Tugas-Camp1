package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func luasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func main() {
	// Masukkan panjang sisi persegi
	sisi := 5.0

	// Hitung luas persegi menggunakan fungsi luasPersegi
	hasilLuas := luasPersegi(sisi)

	// Cetak hasil luas persegi
	fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, hasilLuas)
}
