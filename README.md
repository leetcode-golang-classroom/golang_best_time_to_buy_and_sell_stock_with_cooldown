# golang_best_time_to_buy_and_sell_stock_with_cooldown

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

- After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

**Note:** You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

## Examples

**Example 1:**

```
Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

```

**Example 2:**

```
Input: prices = [1]
Output: 0

```

**Constraints:**

- `1 <= prices.length <= 5000`
- `0 <= prices[i] <= 1000`

## 解析

給定一個 prices 陣列代表 stock 在第 i 天的價格

操作上可以對當日的 prices 做買入或是賣出

但當第 i 天賣出之後，不能第 i + 1 天隔天買入 需要有一個靜止日

同樣的, 當第 i 天買入之後，不能第 i + 1 天隔天買入 需要有一個靜止日

不能連續買，連續賣

要求寫一個演算法算出 給定的 prices 下 能夠做出最大的收益

舉一個例子來做理解

給定 prices: [1, 2, 3, 0, 2]

每天都可以有 買入 或是賣出 的選擇

所以會是以下的決策樹

![](https://i.imgur.com/LY94x6a.png)

可以發現 當有 prices 長度是 n 則共有 $2^n$ 個結果

因為要找出最大的收益所以需要走過所有結點是 O($2^n$)

透過動態規劃可以透過紀錄以計算過的結果避免重複計算 可以簡化時間複雜度為 O(n)

定義 dp[i, buy] 代表第 i 開始買入/賣出 的最大收益 , i 代表 第i 天開始交易, buy 代表該天是否買

dp[i, buy] = max(dp[i+1, false]-prices[i], dp[i+1,true]) if buy == true

dp[i, buy] = max(dp[i+1,false]+prices[i], dp[i+1,false]) if buy == false

為了避免重複計算

需要從後面往前計算

假設從最後一天開始交易

考慮該天買或是賣的最大收益

dp[i, true] = max(dp[i+1, false] - prices[i], dp[i+1, true]) 代表 

第i天買的最大收益 有兩種可能 第i+1天賣的最大收益 收益- 第 i 天價格 或  第i+1天買的最大收益 

第i天賣的最大收益 有兩種可能 第i+2天買的最大收益 收益+ 第i 天價格 或  第i+1天賣的最大收益  


![](https://i.imgur.com/SlMcJRM.png)

而最後所就就是 db_buy[0]

因為一開始沒有股票 第一步只能買

因為只要 loop n 個 開始日 時間複雜度是 O(n)

空間複雜度 也是 O(n)

## 程式碼
```go
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

```
## 困難點

1. 要看出每個開始交易日跟最佳收益之間的遞迴關係

## Solve Point

- [x]  建立兩個長度是 len(prices) +2 陣列用來存放運算過程中的子問題解答 dpSell, dpBuy
- [x]  每次透過遞迴關係式推算每個開始交易的最大收益
- [x]  回傳 dpBuy[0]