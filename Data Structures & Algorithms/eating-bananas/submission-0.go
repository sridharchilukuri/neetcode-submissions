func minEatingSpeed(piles []int, h int) int {

	maxPile := 0
	for _, p := range piles {
		maxPile = max(maxPile, p)
	}

   	left, right := 1, maxPile
	for left < right {
		mid := left + (right-left)/2

		// If this speed works, I'll keep it and try something slower.
		if hoursToEatPiles(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}


func hoursToEatPiles(piles []int, k int) int{
	total:=0

	for _,p := range piles{
         total += (p + k -1)/k
	}

	return total
}
