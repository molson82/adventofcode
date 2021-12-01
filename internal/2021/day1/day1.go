package day1

func Part1(depths []string) int {
	var ans int

	for i, v := range depths {
		if i != 0 {
			if v > depths[i-1] {
				ans++
			}
		}
	}

	return ans
}
