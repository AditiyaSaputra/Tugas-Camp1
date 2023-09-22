package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func luasSegitiga(alas, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}

func main() {
	// Masukkan panjang alas dan tinggi segitiga
	alas := 6.0
	tinggi := 4.0

	// Hitung luas segitiga menggunakan fungsi luasSegitiga
	hasilLuas := luasSegitiga(alas, tinggi)

	// Cetak hasil luas segitiga
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, hasilLuas)
}
