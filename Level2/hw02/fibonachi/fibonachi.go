package fibonachi

// Fibonachi function for calculating the number of fionachi. Takes a uint as input and returns a uint.
func Fibonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibonachi(n-1) + Fibonachi(n-2)
}

// FibonachiCached function for calculating the number of fionachi with cache. Takes a uint as input and returns a uint.
func FibonachiCached(num uint) uint {

	fibonachiNumbers := make(map[uint]uint)

	if _, ok := fibonachiNumbers[num]; ok {
		return fibonachiNumbers[num]
	} else {
		fibonachiNumbers[num] = Fibonachi(num)

		return fibonachiNumbers[num]
	}
}
