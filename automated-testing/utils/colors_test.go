package utils

import (
	"reflect"
	"testing"
)

func TestCountColors(t *testing.T) {
	colors := Colors{
		"red":   true,
		"green": true,
		"blue":  true,
	}

	cases := []struct {
		name           string
		colors         Colors
		text           string
		expectedOutput ColorCounts
	}{
		{
			name:           "empty string",
			colors:         colors,
			text:           "",
			expectedOutput: ColorCounts{},
		},
		{
			name:   "valid input",
			colors: colors,
			text:   "My favorite colors are red, green, blue.",
			expectedOutput: ColorCounts{
				"red":   1,
				"green": 1,
				"blue":  1,
			},
		},
		{
			name:   "consecutive colors",
			colors: colors,
			text:   "My favorite colors are redred, green, blue.",
			expectedOutput: ColorCounts{
				"green": 1,
				"blue":  1,
			},
		},
	}

	for _, c := range cases {
		if output := CountColors(c.colors, c.text); !reflect.DeepEqual(output, c.expectedOutput) {
			t.Errorf("\ncase: %s\ninput: %v, %s\ngot: %v\nwant: %v", c.name, c.colors, c.text, output, c.expectedOutput)
		}
	}
}
