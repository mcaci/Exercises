package check

func Fermat(numberToTest uint) bool {
	var base uint = 2
	var tryNumber uint = 1
	const maxTryNumber = 20

	var isPrime bool = true
	for base <= numberToTest && tryNumber < maxTryNumber && isPrime {
		isPrime = ((base)^(numberToTest-1) - 1) == 1
		base++
		tryNumber++
	}
	return true
}