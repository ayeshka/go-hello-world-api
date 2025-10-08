package validation

import (
	"strings"
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid name starting with A", "Alice", true},
		{"Valid name starting with a (lowercase)", "alice", true},
		{"Valid name starting with B", "Bob", true},
		{"Valid name starting with b (lowercase)", "bob", true},
		{"Valid name starting with C", "Charlie", true},
		{"Valid name starting with D", "David", true},
		{"Valid name starting with E", "Emma", true},
		{"Valid name starting with F", "Frank", true},
		{"Valid name starting with G", "Grace", true},
		{"Valid name starting with H", "Henry", true},
		{"Valid name starting with I", "Isabella", true},
		{"Valid name starting with J", "John", true},
		{"Valid name starting with K", "Kate", true},
		{"Valid name starting with L", "Liam", true},
		{"Valid name starting with M", "Mark", true},
		{"Valid name starting with m (lowercase)", "mary", true},
		{"Valid single character A", "A", true},
		{"Valid single character a", "a", true},
		{"Valid single character M", "M", true},
		{"Valid single character m", "m", true},
		{"Valid name with spaces", "John Doe", true},
		{"Valid name Alice Nam", "Alice Nam", true},
		{"Valid lowercase alice nam", "alice nam", true},
		{"Valid name with hyphen", "Anne-Marie", true},
		{"Valid multi-word with middle initial", "Mary J Watson", true},

		// Invalid names (N-Z)
		{"Invalid name starting with N", "Nick", false},
		{"Invalid name starting with n (lowercase)", "nick", false},
		{"Invalid name starting with O", "Oliver", false},
		{"Invalid name starting with P", "Peter", false},
		{"Invalid name starting with Q", "Quinn", false},
		{"Invalid name starting with R", "Rachel", false},
		{"Invalid name starting with S", "Sarah", false},
		{"Invalid name starting with T", "Tom", false},
		{"Invalid name starting with U", "Uma", false},
		{"Invalid name starting with V", "Victor", false},
		{"Invalid name starting with W", "William", false},
		{"Invalid name starting with X", "Xavier", false},
		{"Invalid name starting with Y", "Yara", false},
		{"Invalid name starting with Z", "Zane", false},
		{"Invalid name starting with z (lowercase)", "zoe", false},
		{"Invalid single character N", "N", false},
		{"Invalid single character n", "n", false},
		{"Invalid single character Z", "Z", false},
		{"Invalid single character z", "z", false},
		{"Invalid multi-word starting with N", "Nancy Smith", false},
		{"Invalid multi-word starting with Z", "Zoe Johnson", false},
		{"Invalid lowercase multi-word starting with n", "nancy smith", false},

		// Edge cases
		{"Empty string", "", false},
		{"Name starting with number", "123John", false},
		{"Name starting with special character @", "@Alice", false},
		{"Name starting with special character #", "#Bob", false},
		{"Name starting with special character $", "$Charlie", false},
		{"Name starting with special character %", "%David", false},
		{"Name starting with underscore", "_Emma", false},
		{"Name starting with hyphen", "-Frank", false},
		{"Name starting with space", " Grace", false},
		{"Name starting with tab", "\tHenry", false},
		{"Name starting with newline", "\nIsabella", false},

		// Unicode edge cases (accented characters are treated as separate characters)
		{"Name with accented A", "Àlex", false},
		{"Name with accented E", "Émilie", false},
		{"Name with accented O", "Óliver", false},

		// Boundary cases
		{"Boundary case - exactly M", "Michael", true},
		{"Boundary case - exactly N", "Nathan", false},
		{"Boundary case - exactly m", "michelle", true},
		{"Boundary case - exactly n", "nancy", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateName(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateName(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateNameBoundaryConditions(t *testing.T) {
	// Test all letters at the boundary
	validLetters := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M'}
	invalidLetters := []rune{'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	// Test valid letters (A-M)
	for _, letter := range validLetters {
		name := string(letter)
		if !ValidateName(name) {
			t.Errorf("ValidateName(%q) should return true for letter in A-M range", name)
		}

		// Test lowercase version
		lowerName := strings.ToLower(name)
		if !ValidateName(lowerName) {
			t.Errorf("ValidateName(%q) should return true for lowercase letter in a-m range", lowerName)
		}
	}

	// Test invalid letters (N-Z)
	for _, letter := range invalidLetters {
		name := string(letter)
		if ValidateName(name) {
			t.Errorf("ValidateName(%q) should return false for letter in N-Z range", name)
		}

		// Test lowercase version
		lowerName := strings.ToLower(name)
		if ValidateName(lowerName) {
			t.Errorf("ValidateName(%q) should return false for lowercase letter in n-z range", lowerName)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{"Letter A", 'A', true},
		{"Letter z", 'z', true},
		{"Number 1", '1', false},
		{"Special char @", '@', false},
		{"Space", ' ', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLetter(tt.input)
			if result != tt.expected {
				t.Errorf("IsLetter(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
