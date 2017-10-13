package utils

import (
	"fmt"
	"regexp"
)

type Colors map[string]bool
type ColorCounts map[string]int

// CountColors function counts how many pre-defined colors
// appear in a given text.
func CountColors(colors Colors, text string) ColorCounts {

	colorsCounts := make(ColorCounts)

	r := regexp.MustCompile("\\b\\w+\\b")
	words := r.FindAllString(text, -1)
	fmt.Println(words)
	for _, w := range words {
		if _, isColor := colors[w]; isColor {
			colorsCounts[w]++
		}
	}

	return colorsCounts
}
