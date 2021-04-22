
// @Title: 买卖股票的最佳时机 (Best Time to Buy and Sell Stock)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 18:23:42
// @Runtime: 240 ms
// @Memory: 16.9 MB

func maxProfit(prices []int) int {
    //f[i][0]表示第i天不持有股票的收益
    //f[i][1]表示第i天持有股票的收益
    f:=make([][]int,len(prices))
    for i:=0;i<len(prices);i++{
        f[i]=make([]int,2)
    }
    f[0][0]=0
    f[0][1]=-prices[0]
    for i:=1;i<len(prices);i++{
        f[i][0]=max(f[i-1][0],f[i-1][1]+prices[i])
        f[i][1]=max(f[i-1][1],-prices[i])
    }
    return f[len(prices)-1][0]

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
