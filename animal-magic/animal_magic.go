package chance

import (
	"math/rand"
	"time"
)

var r1 *rand.Rand

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	return rollADiceWithMax(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return float64(r1.Intn(120)) / 10
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

// returns a random int d with 0 <= d < value
func rollADiceWithMax(value int) int {
	return r1.Intn(value)
}
