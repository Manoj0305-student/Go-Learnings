package arrays
//{1, 7, 0, 7, 7, 2, 1, 7}
func GoldStarTelescope(telescope []int, k int) int {

	currentCount := 0

	goldStar := 7

	for i:=0;i<k;i++ {
		if telescope[i] == goldStar {
			currentCount++
		}
	}

	maxCount := currentCount

	for i:=k;i<len(telescope);i++ {
		if telescope[i] == goldStar {
			currentCount++
		}
		if telescope[i-k] == goldStar {
			currentCount--
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
} 