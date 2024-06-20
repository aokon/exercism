package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

func formatTable(table int) string {
	var result string

	switch {
	case table < 10:
		result = fmt.Sprintf("00%d", table)
	case table >= 10 && table < 100:
		result = fmt.Sprintf("0%d", table)
	default:
		result = fmt.Sprintf("%d", table)
	}

	return result
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\n", name) +
		fmt.Sprintf("You have been assigned to table %s. ", formatTable(table)) +
		fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance) +
		fmt.Sprintf("You will be sitting next to %s.", neighbor)
}
