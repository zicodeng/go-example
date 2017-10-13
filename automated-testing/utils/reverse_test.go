package utils

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			"empty string",
			"",
			"",
		},
		{
			"single character",
			"a",
			"a",
		},
		{
			"two characters",
			"ab",
			"ba",
		},
		{
			"odd number of characters",
			"abc",
			"cba",
		},
		{
			"even number of characters",
			"abcd",
			"dcba",
		},
		{
			"palindrome",
			"aibohphobia",
			"aibohphobia",
		},
		{
			"non-ASCII character",
			"Hello, 世界",
			"界世 ,olleH",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectedOutput {
			t.Errorf("\ncase: %s\ninput: %s\ngot: %v\nwant: %s", c.name, c.input, output, c.expectedOutput)
		}
	}
}
