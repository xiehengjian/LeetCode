
// @Title: 零钱兑换 (Coin Change)
// @Author: 846188037@qq.com
// @Date: 2021-04-13 21:21:52
// @Runtime: 12 ms
// @Memory: 6.3 MB


func coinChange(coins []int, amount int) int {
    // 状态 dp[i]表示金额为i时，组成的最小硬币个数
    // 推导 dp[i]  = min(dp[i-1], dp[i-2], dp[i-5])+1, 前提 i-coins[j] > 0
    // 初始化为最大值 dp[i]=amount+1
    // 返回值 dp[n] or dp[n]>amount =>-1
    dp:=make([]int,amount+1)

	//初始化最大值
    for i:=0;i<=amount;i++{
        dp[i]=amount+1
    }

	//当金额为0时有0种方法
    dp[0]=0

	//开始递推
    for i:=1;i<=amount;i++{//算每个金额
        for j:=0;j<len(coins);j++{//每种选择
            if  i-coins[j]>=0  {
                dp[i]=min(dp[i],dp[i-coins[j]]+1)
            }
        }
    }

	//返回结果
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]

}
func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}



