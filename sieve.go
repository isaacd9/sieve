package sieve

// FromTo calculates the list of primes between start and end
func FromTo(start, end uint32) (ret []uint32) {
	arr := make([]bool, end*end)

	var i, x uint32
	for i = start; i < end; i++ {
		if arr[i] == false {
			ret = append(ret, i)
		}

		for x = 0; x <= i; x++ {
			arr[x*i] = true
		}
	}

	return ret
}

// To calculates the list of primes until end
func To(end uint32) (ret []uint32) {
	return FromTo(2, end)
}
