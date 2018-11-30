package bloomfilter

import "testing"

func TestAddBloomFilter(t *testing.T) {
	bloomfilter := initBloomFilterForTest(100)

	added := false
	for _, bit := range bloomfilter.bitarray {
		added = bit || added
		if added {
			break
		}
	}
	if !added {
		t.Fatal("Nothing was added")
	}
}

func TestStringIsNotThereForSure(t *testing.T) {
	bloomfilter := initBloomFilterForTest(100)
	testString := "I'm-not-there"
	testValue := bloomfilter.IsPossiblyThere(testString)
	if testValue {
		t.Fatalf("\"%s\" should never be present!", testString)
	}
}

func TestStringMayBePresent(t *testing.T) {
	bloomfilter := initBloomFilterForTest(100)
	testString := "test"
	testValue := bloomfilter.IsPossiblyThere(testString)
	if !testValue {
		t.Fatalf("\"%s\" may be present but it hasn't been found", testString)
	}
}

func initBloomFilterForTest(size int) *BloomFilter {
	var bloomfilter BloomFilter
	bloomfilter.Init(size)
	bloomfilter.Add("test")
	return &bloomfilter
}
