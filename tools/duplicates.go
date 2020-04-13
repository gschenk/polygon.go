// Package tools provides useful functions for PolyGo
package tools

import (
	"polyGo/vector"
)

// find duplicate vectors in slice and return slice of unique values
// based on Gist by arehmandev
// https://gist.github.com/arehmandev/4790544bf08f9965596eb0d75f9f270b

func RemoveDuplicates(elements []vector.Vec) []vector.Vec {
	// Use map to record duplicates as we find them.
	encountered := map[vector.Vec]bool{}
	result := []vector.Vec{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// usage: RemoveDuplicates(slice)
