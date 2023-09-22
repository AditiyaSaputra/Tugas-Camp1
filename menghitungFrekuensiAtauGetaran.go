package main

import (
	"fmt"
)

func main() {
	// Nilai periode yang sudah tersedia
	periode := 0.5 // nilai periode yang diinginkan (dalam detik)

	// Hitung frekuensi
	frekuensi := 1 / periode

	// Cetak hasil frekuensi
	fmt.Printf("Frekuensi: %.2f Hz\n", frekuensi)
}
