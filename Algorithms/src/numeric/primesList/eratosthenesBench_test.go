package primesList

import (
	"testing"
)

func BenchmarkListOfPrimesUpTo100_optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErathosthenesSieve(5000)
	}
}

func BenchmarkListOfPrimesUpTo100_unoptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErathosthenesSieveUnoptimized(5000)
	}
}

func ErathosthenesSieveUnoptimized(limit uint) *[]uint {
	isNotPrimeList := make([]bool, limit+1)
	var number uint
	for number = 2; number <= limit; number++ {
		for i := uint(2); number * i < uint(len(isNotPrimeList)); i++ {
			isNotPrimeList[number * i] = true
		}
	}
	
	var list []uint
	for i, isNotPrime := range isNotPrimeList[2:] {
		if !isNotPrime {
			list = append(list, uint(i+2))
		}
	}
	return &list
}