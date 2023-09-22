package main

import (
	"fmt"
)

func main() {
	// Nilai periode yang sudah tersedia
	periode := 0.5 // Ganti dengan nilai periode yang Anda inginkan (dalam detik)

	// Hitung frekuensi
	frekuensi := 1 / periode

	// Cetak hasil frekuensi
	fmt.Printf("Frekuensi: %.2f Hz\n", frekuensi)
}
