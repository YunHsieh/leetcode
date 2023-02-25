/*
Time complexity: O(N)
Space complexity: O(1)

Runtime Faster than 99%+
*/

func maxProfit(prices []int) int {
    result := 0
    lastData := prices[0]
    tmp := 0
    for i := 0; len(prices) > i; i++ {
        if lastData > prices[i]{
            lastData = prices[i]
        } else {
            tmp = prices[i] - lastData
            if result < tmp {
                result = tmp
            }
        }
    }
    return result
}
