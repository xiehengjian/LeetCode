
// @Title: 青蛙跳台阶问题 (青蛙跳台阶问题  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 09:22:25
// @Runtime: 0 ms
// @Memory: 1.9 MB

func numWays(n int) int {
    //f(n)=f(n-1)+f(n-2)
    if n==0 || n==1{
        return 1
    }
    if n==2{
        return 2
    }
    a1,a2,a3:=1,2,0
    for i:=3;i<n+1;i++{
        a3=(a1+a2)%1000000007
        a1=a2
        a2=a3
    }
    return a3

}
