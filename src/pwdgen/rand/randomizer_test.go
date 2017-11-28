package rand

import (
	"testing"
)

// Runs RandomRune a number of times, returns a map with the number
// of occurances of each rune in the given charset string.
func SampleRandomRunes(t *testing.T, charset string, numsamples int) map[rune]int {
	results := make(map[rune]int)
	for _, v := range charset {
		results[v] = 0
	}

	var char rune

	// run the sampling
	for i := 0; i < numsamples; i++ {
		char = RandomRune(charset)
		//char = fn(charset)

		// count
		if _, ok := results[char]; ok {
			results[char]++
		} else {
			t.Errorf("in iteration %d: got value \"%#v\" which is not in the charset",
				i, char)
		}

	}

	return results
}

// When we ask for a random rune, we want:
// a) get a rune exclusively from the charset
// b) after many calls, each char from the charset must be returned
// We consider 1000 calls for a charset of 10 characters as many.
func TestRandomRune(t *testing.T) {
	const numsamples = 1000
	const charsetstring = "0123456789"

	results := SampleRandomRunes(t, charsetstring, numsamples)

	// check the result map
	for key, value := range results {
		if value == 0 {
			t.Errorf("character %#v was never picked", key)
		}
	}
}

// --- BENCHMARKS ---------------------------------------------------

func BenchmarkDefaultRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultRandom(100)
	}
}
