package logs

import "unicode/utf8"

const recommendation rune = '❗'
const search rune = '🔍'
const weather rune = '☀'

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        if char == recommendation {
            return "recommendation"
        } else if char == search {
        	return "search"
        } else if char == weather {
        	return "weather"
        }
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newString string = ""
    for _, char := range log {
        if char == oldRune {
            newString += string(newRune)
        } else {
        	newString += string(char)
        }
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := utf8.RuneCountInString(log)
    return length <= limit
}