package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToLower(customer)
	customer = strings.ToTitle(customer)
	return "Welcome to the Tech Palace, " + customer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.Trim(oldMsg, "*\n")
	msg = strings.TrimSpace(msg)
	return msg
}
