package sol

func maxProfit(prices []int) int {
	pLen := len(prices)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// create dpBuy, dpSell for store processed result
	// dpBuy[i] denote start transaction from i with buy
	// dpSell[i] denote start transaction from i with buy
	// dpBuy[i] = max(dpBuy[i+1], dpSell[i+1] - prices[i])
	// dpSell[i] = max(dpBuy[i+2]+prices[i], dbSell[i+1]), because could not buy after sell the next day
	dpBuy, dpSell := make([]int, pLen+2), make([]int, pLen+2)
	for start := pLen - 1; start >= 0; start-- {
		dpBuy[start] = max(dpBuy[start+1], dpSell[start+1]-prices[start])
		dpSell[start] = max(dpSell[start+1], dpBuy[start+2]+prices[start])
	}
	return dpBuy[0]
}
