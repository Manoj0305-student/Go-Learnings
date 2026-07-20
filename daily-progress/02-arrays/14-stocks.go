package arrays 

/* 
 MaxProfit calculates the maximum profit possible from buying and selling a stock 
 given an array of daily prices. You must buy before you sell.

 Supported Test Cases:

 Case 1: Standard profitable transaction
 Input Array: [7, 1, 5, 3, 6, 4]
 Expected Output: 5 (Buy on day 2 at price 1, sell on day 5 at price 6)

 Case 2: Prices continuously decrease (No profit possible)
 Input Array: [7, 6, 4, 3, 1]
 Expected Output: 0

 Case 3: Best buying day is the first day, best selling day is the last day
 Input Array: [1, 2, 9]
 Expected Output: 8
*/
func Stocks(prices []int) int {

	maxProfit := 0
	bestBuy := prices[0]

	for i:=0 ; i < len(prices) ; i++ {
		if prices[i] < bestBuy {
			bestBuy = prices[i]
		}

		profit := prices[i] - bestBuy

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}