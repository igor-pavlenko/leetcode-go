package leetcode_0121_Best_Time_to_Buy_and_Sell_Stock

// Brute Force
func maxProfitBF(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if (prices[j] - prices[i]) > max {
				max = prices[j] - prices[i]
			}
		}
	}

	return max
}

func maxProfit(prices []int) int {
	minprice := 1<<31 - 1
	maxprofit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}

	return maxprofit
}
