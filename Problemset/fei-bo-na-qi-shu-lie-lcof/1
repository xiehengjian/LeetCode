
// @Title: 斐波那契数列 (斐波那契数列  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 09:19:40
// @Runtime: 0 ms
// @Memory: 1.9 MB

func fib(n int) int {
    if n==0{
        
        return 0
    }
    if n==1{
        return 1
    }
    a1,a2,a3:=0,1,0
    for i:=2;i<n+1;i++{
        a3=(a1+a2)%1000000007
        a1=a2 
        a2=a3
    }
    return a3%1000000007

}
