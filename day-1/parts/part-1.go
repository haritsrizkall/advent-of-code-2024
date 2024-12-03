package parts

func PartOne(leftPair, rightPair []int) int {
	res := 0
	for i := 0; i < len(leftPair); i++ {
		temp := leftPair[i] - rightPair[i]
		if temp < 0 {
			temp = temp * (-1)
		}
		res += temp
	}

	return res
}
