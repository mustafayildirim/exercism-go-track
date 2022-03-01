package techpalace

import (
	"bytes"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcomeMessage := "Welcome to the Tech Palace, "
	return welcomeMessage + strings.ToUpper(customer)
}

func getMessageBorder(numStarsPerLine int) string {
	var buffer bytes.Buffer

	for i := 0; i < numStarsPerLine; i++ {
		buffer.WriteString("*")
	}

	return buffer.String()
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	border := getMessageBorder(numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	oldMsg = strings.Replace(oldMsg, "*", "", -1)
	oldMsg = strings.Replace(oldMsg, "\n", "", -1)
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}
