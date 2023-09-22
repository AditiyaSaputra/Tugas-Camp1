package main

import (
	"fmt"
	"math"
)

func main() {
	// Nilai jari-jari lingkaran yang sudah tersedia
	jariJari := 5.0

	// Hitung luas lingkaran
	luas := math.Pi * math.Pow(jariJari, 2)

	// Cetak hasil luas lingkaran
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luas)
}
