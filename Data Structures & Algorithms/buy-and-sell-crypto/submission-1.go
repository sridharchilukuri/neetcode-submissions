func maxProfit(prices []int) int {

	maxProfit := 0

	i, j:= 0,1

	for i < j && j < len(prices){

        profit := prices[j] - prices[i]

		if profit < maxProfit {
			if(prices[i]> prices[j]){
				i = j
				j++
			}else{
				j++
			}
		}else{
			maxProfit = profit
			j++
		}
	}

	return maxProfit

}
