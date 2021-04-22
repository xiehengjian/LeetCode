
// @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
// @Author: 846188037@qq.com
// @Date: 2020-07-16 20:44:28
// @Runtime: 0 ms
// @Memory: 2.3 MB

func kidsWithCandies(candies []int, extraCandies int) []bool {
    n := len(candies)
    maxCandies := 0
    for i := 0; i < n; i++ {
        maxCandies = max(maxCandies, candies[i])
    }
    ret := make([]bool, n)
    for i := 0; i < n; i++ {
        ret[i] = candies[i] + extraCandies >= maxCandies
    }
    return ret
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}


