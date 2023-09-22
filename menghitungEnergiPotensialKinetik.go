package main

import (
	"fmt"
)

func main() {
	// Nilai massa, tinggi, dan kecepatan yang sudah tersedia
	massa := 10.0     // nilai massa yang Anda inginkan (dalam kilogram)
	gravitasi := 9.81 // Percepatan gravitasi bumi (dalam m/s^2)
	tinggi := 2.0     // nilai tinggi yang Anda inginkan (dalam meter)
	kecepatan := 5.0  // nilai kecepatan yang Anda inginkan (dalam m/s)

	// Hitung energi potensial
	energiPotensial := massa * gravitasi * tinggi

	// Hitung energi kinetik
	energiKinetik := 0.5 * massa * (kecepatan * kecepatan)

	// Cetak hasil energi potensial dan kinetik
	fmt.Printf("Energi Potensial: %.2f Joule\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f Joule\n", energiKinetik)
}
