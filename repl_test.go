package main

import (
	"testing"
)



func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: 	"ChArManDer SnORlax    ",
			expected: []string{"charmander", "snorlax"},
		},
		{
			input: " GENGar ",
			expected: []string{"gengar"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)

		if actualLen != expectedLen {
			t.Errorf("Error! Actual length and expected length do not match: Actual Len %v, Expected Len %v", actualLen, expectedLen)
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error! Actual word and expected word do not match: Actual Word %v, Expected Word %v", word, expectedWord)
				t.Fail()
			}
		}
	}

}