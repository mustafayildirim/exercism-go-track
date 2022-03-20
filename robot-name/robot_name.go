package robotname

import (
	"errors"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

const max = 26 * 26 * 10 * 10 * 10

var used int = 0
var names = make(map[string]bool, 0)

func init() {
	rand.Seed(time.Now().Unix())
}

func (r *Robot) Name() (string, error) {

	if r.name == "" {
		if len(names) >= max {
			return "", errors.New("No names left!")
		}

		r.name = newName()
	}

	return r.name, nil
}

func newName() string {
	name := generateName()
	exists := true
	for exists && used < max {
		name = generateName()
		_, exists = names[name]
	}
	names[name] = true
	used++
	return name
}

func generateName() string {
	name := RandStringRunes(2)
	name = name + RandNumberRunes(3)

	return name

}

func (r *Robot) Reset() {
	r.name = ""
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumberRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}
