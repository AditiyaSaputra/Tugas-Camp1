package main

import (
	"fmt"
)

func main() {
	// Nilai suhu dalam Celsius yang sudah tersedia
	celsius := 25.0 // Ganti dengan nilai Celsius yang Anda inginkan

	// Konversi suhu ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Konversi suhu ke Kelvin
	kelvin := celsius + 273.15

	// Cetak hasil konversi
	fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, fahrenheit)
	fmt.Printf("%.2f Celsius = %.2f Kelvin\n", celsius, kelvin)

	// Nilai suhu dalam Fahrenheit yang sudah tersedia
	fahrenheit = 77.0 // Ganti dengan nilai Fahrenheit yang Anda inginkan

	// Konversi suhu ke Celsius
	celsius = (fahrenheit - 32) * 5 / 9

	// Konversi suhu ke Kelvin
	kelvin = (fahrenheit-32)*5/9 + 273.15

	// Cetak hasil konversi
	fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", fahrenheit, celsius)
	fmt.Printf("%.2f Fahrenheit = %.2f Kelvin\n", fahrenheit, kelvin)

	// Nilai suhu dalam Kelvin yang sudah tersedia
	kelvin = 298.15 // Ganti dengan nilai Kelvin yang Anda inginkan

	// Konversi suhu ke Celsius
	celsius = kelvin - 273.15

	// Konversi suhu ke Fahrenheit
	fahrenheit = (celsius * 9 / 5) + 32

	// Cetak hasil konversi
	fmt.Printf("%.2f Kelvin = %.2f Celsius\n", kelvin, celsius)
	fmt.Printf("%.2f Kelvin = %.2f Fahrenheit\n", kelvin, fahrenheit)
}
