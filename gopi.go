package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var inside, total, pi float64
	rand.Seed(time.Now().Unix())
	for total < 10000 {
		x := rand.Float64()
		y := rand.Float64()
		distance := math.Pow(x, 2) + math.Pow(y, 2)
		if distance < 1 {
			inside = inside + 1
		}
		total = total + 1
	}
	pi = float64(4 * (inside / total))
	fmt.Println(pi)
}
