// Package proverb implements a simple library that generates a proverbial rhyme based on the input slice of strings
package proverb

import "fmt"

// Proverb returns the proverbial rhyme based on the input slice of strings
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	var output []string
	for i := 0; i < len(rhyme)-1; i++ {
		output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return output
}
