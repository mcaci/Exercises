package primesList

// import "fmt"

func ErathosthenesSieve(lastNumberToCheck uint) *[]uint {
	isNotPrimeList := createIsNotPrimeList(lastNumberToCheck)
	fillIsNotPrimeList(&isNotPrimeList, lastNumberToCheck)
	return buildPrimesList(&isNotPrimeList)
}

func createIsNotPrimeList(lastNumberToCheck uint) []bool {
	const sizeFor0and1 = 2
	const sizeForLastNumberToCheck = 1
	return make([]bool, lastNumberToCheck + sizeForLastNumberToCheck - sizeFor0and1)
}

func fillIsNotPrimeList(isNotPrimeList *[]bool, lastNumberToCheck uint) {
	for possiblePrime := 2; possiblePrime * possiblePrime <= int(lastNumberToCheck); possiblePrime++ {
		for factor := possiblePrime; (possiblePrime * factor - 2) < len(*isNotPrimeList); factor++ {
			(*isNotPrimeList)[(possiblePrime * factor) - 2] = true
		}
	}
}

func buildPrimesList(isNotPrimeList *[]bool) *[]uint {
	var list []uint
	for i, isNotPrime := range (*isNotPrimeList) {
		if !isNotPrime {
			list = append(list, uint(i+2))
		}
	}
	return &list
}