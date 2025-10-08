package validation

import (
	"strings"
	"unicode"
)

// ValidateName validates if a name is valid based on the first letter rule.
// Returns true if name starts with A-M (case-insensitive), false otherwise.
// Supports both single names and multi-word names (e.g., "Alice", "Alice Nam", "John Doe Jr").
//
// Rules:
// - Name must not be empty
// - First character must be a letter
// - First letter must be in range A-M (case-insensitive)
// - Multi-word names are validated based only on the first character of the entire string
func ValidateName(name string) bool {
	// Check if name is missing or empty
	if name == "" {
		return false
	}

	// Get the first character of the name
	firstChar := rune(strings.ToUpper(name)[0])

	// Check if the first character is a letter
	if !unicode.IsLetter(firstChar) {
		return false
	}

	// Check if first letter is in the first half of alphabet (A-M)
	return firstChar >= 'A' && firstChar <= 'M'
}

// IsLetter checks if a character is a letter
func IsLetter(char rune) bool {
	return unicode.IsLetter(char)
}
