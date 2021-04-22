
// @Title: 买卖股票的最佳时机 II (Best Time to Buy and Sell Stock II)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 21:09:13
// @Runtime: 4 ms
// @Memory: 4.1 MB

func maxProfit(prices []int) int {
    f:=make([][]int,len(prices))
    for i:=0;i<len(prices);i++{
        f[i]=make([]int,2)
    }
    f[0][0]=0
    f[0][1]=-prices[0]
    for i:=1;i<len(prices);i++{
        f[i][0]=max(f[i-1][0],f[i-1][1]+prices[i])
        f[i][1]=max(f[i-1][1],f[i-1][0]-prices[i])
    }
    return f[len(prices)-1][0]

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
