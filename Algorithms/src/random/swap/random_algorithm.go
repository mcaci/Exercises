package swap

import (
	"../lcg"
)

func Rand(array []int) []int {
	var random lcg.LCG
	max := len(array)
	var copyArray = make([]int, max)
	copy(copyArray[:], array)
	for i := range copyArray {
		// pick random index between i and size
		// careful on this part (must be correct random number)
		random.Init(max+3*i, (max-i)*2+1, i, max-i)
		j := random.Next() + i
		// swap i and j
		temp := copyArray[i]
		copyArray[i] = copyArray[j]
		copyArray[j] = temp
	}
	return copyArray
}
