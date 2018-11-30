package bloomfilter

import (
	"hash"
	"hash/crc32"
	"hash/fnv"
)

type BloomFilter struct {
	bitarray []bool
}

func (bloomfilter *BloomFilter) Init(arraySize int) {
	bloomfilter.bitarray = make([]bool, arraySize)
}

func (bloomfilter *BloomFilter) Add(s string) {
	index1 := hashToIndex(s, uint32(len(bloomfilter.bitarray)), fnv.New32a())
	index2 := hashToIndex(s, uint32(len(bloomfilter.bitarray)), crc32.NewIEEE())
	bloomfilter.bitarray[index1] = true
	bloomfilter.bitarray[index2] = true
}

func (bloomfilter *BloomFilter) IsPossiblyThere(s string) bool {
	index1 := hashToIndex(s, uint32(len(bloomfilter.bitarray)), fnv.New32a())
	index2 := hashToIndex(s, uint32(len(bloomfilter.bitarray)), crc32.NewIEEE())
	return bloomfilter.bitarray[index1] && bloomfilter.bitarray[index2]
}

func hashToIndex(s string, arraySize uint32, hashMethod hash.Hash32) uint32 {
	hashMethod.Write([]byte(s))
	hash := hashMethod.Sum32()
	return hash % arraySize
}
