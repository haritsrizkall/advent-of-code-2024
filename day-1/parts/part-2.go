package parts

func PartTwo(leftPair, rightPair []int) int {
	freqsMap := make(map[int]int)
	for _, val := range rightPair {
		_, ok := freqsMap[val]
		if !ok {
			freqsMap[val] = 1
		} else {
			freqsMap[val] += 1
		}
	}

	res := 0

	for _, val := range leftPair {
		freq, ok := freqsMap[val]
		if ok {
			res = res + (val * freq)
		}
	}

	return res
}
