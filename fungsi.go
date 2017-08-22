package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var names = []string{"John", "Wick"}

/*
  rand.Seed(time.Now().Unix())
  memastikan bahwa angka random yang akan di-generate benar-benar acak. Kita bisa gunakan angka apa saja sebagai nilai parameter fungsi ini (umumnya diisi time.Now().Unix()).

  FUNGSI
  - return value
    func nama(params tipeparam) tipereturn { return value}
  - declare param tipe data sama
    func nama(param1, param2, param3 tipe) tipereturn
  - declare param beda beda
    func nama(param1 tipe1, param2 tipe2)
*/

func main() {

	// printMessage("halo", names)
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
