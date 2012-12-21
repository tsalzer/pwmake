package stats

import (
    "math"
)

// Basic (and simple) entropy calculation.
func EntropyOf(numSymbols int, length int) float64 {
    h := float64(length) * (math.Log(float64(numSymbols)) / math.Log(2))
    return h
}

// calculate the entropy for a given character set.
// This is the strength of a password of the given length generated
// randomly from the given character set.
func GetEntropy(charset string, length int) float64 {
    return EntropyOf(len(charset), length)
}

