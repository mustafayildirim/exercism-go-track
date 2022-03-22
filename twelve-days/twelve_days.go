package twelve

import "fmt"

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var songPart = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

func getSongPart(i int) string {
	if i == 1 {
		return songPart[1]
	}

	songLine := ""
	for i > 1 {
		songLine += fmt.Sprintf("%s, ", songPart[i])
		i--
	}

	songLine += fmt.Sprintf("and %s", songPart[1])

	return songLine
}

func Verse(i int) string {
	line := getSongPart(i)
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[i], line)
}

func Song() string {
	song := ""
	for i := 1; i <= 11; i++ {
		song += fmt.Sprintf("%s\n", Verse(i))
	}
	song += fmt.Sprintf("%s", Verse(12))
	return song
}
