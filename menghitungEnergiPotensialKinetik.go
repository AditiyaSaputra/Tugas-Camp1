package main

import (
	"fmt"
)

func main() {
	// Nilai massa, tinggi, dan kecepatan yang sudah tersedia
	massa := 10.0     // nilai massa yang diinginkan (dalam kilogram)
	gravitasi := 9.81 // Percepatan gravitasi bumi (dalam m/s^2)
	tinggi := 2.0     // nilai tinggi yang diinginkan (dalam meter)
	kecepatan := 5.0  // nilai kecepatan yang diinginkan (dalam m/s)

	// Hitung energi potensial
	energiPotensial := massa * gravitasi * tinggi

	// Hitung energi kinetik
	energiKinetik := 0.5 * massa * (kecepatan * kecepatan)

	// Cetak hasil energi potensial dan kinetik
	fmt.Printf("Energi Potensial: %.2f Joule\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f Joule\n", energiKinetik)
}
