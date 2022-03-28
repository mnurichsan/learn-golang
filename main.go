package main

import (
	"belajar-golang/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello Belajar Golang")

	result := calculation.Hitung(80, 80, 80)
	fmt.Println(result)
	hasilAkhir := calculation.Score(result)

	fmt.Println(hasilAkhir)

}
