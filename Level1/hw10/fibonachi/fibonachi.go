package fibonachi

func Fibonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonachi(n-1) + Fibonachi(n-2)
}

func FibonachiCached(num uint) uint {

	fibonachiNumbers := make(map[uint]uint)

	if _, ok := fibonachiNumbers[num]; ok {
		return fibonachiNumbers[num]
	} else {
		fibonachiNumbers[num] = Fibonachi(num)
		return fibonachiNumbers[num]
	}
}
