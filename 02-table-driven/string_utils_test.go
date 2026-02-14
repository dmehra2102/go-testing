package tabledriven

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "simple palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "palindrome with spaces",
			input:    "A man a plan a canal Panama",
			expected: true,
		},
		{
			name:     "not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "mixed case palindrome",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "palindrome with punctuation",
			input:    "Was it a car or a cat I saw",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "Hello, 世界", "界世 ,olleH"},
		{"numbers", "12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"no vowels", "bcd", 0},
		{"all vowels", "aeiou", 5},
		{"mixed case", "Hello World", 3},
		{"empty string", "", 0},
		{"uppercase", "AEIOU", 5},
		{"sentence", "The quick brown fox", 5},
		{"single vowel", "a", 1},
		{"consonants only", "xyz", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringEdgeCases(t *testing.T) {
	t.Run("empty strings", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func(string) any
			want any
		}{
			{"IsPalindrome", func(s string) any { return IsPalindrome(s) }, true},
			{"Reverse", func(s string) any { return Reverse(s) }, ""},
			{"CountVowels", func(s string) any { return CountVowels(s) }, 0},
			{"WordCount", func(s string) any { return WordCount(s) }, 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.fn("")
				if result != tt.want {
					t.Errorf("%s(\"\") = %v; want %v", tt.name, result, tt.want)
				}
			})
		}
	})
}
