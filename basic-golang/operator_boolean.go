package main

import "fmt"

func main() {
	var (
		nilaiAkhir = 80
		Absensi    = 75

		lulusNilaiAkhir = nilaiAkhir > 80
		lulusAbsensi    = Absensi > 80
	)
	fmt.Println(lulusNilaiAkhir && lulusAbsensi)
	fmt.Println(lulusNilaiAkhir || lulusAbsensi)
	fmt.Println(!lulusNilaiAkhir)
}
