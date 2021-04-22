
// @Title: 爬楼梯 (Climbing Stairs)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 13:26:58
// @Runtime: 0 ms
// @Memory: 1.9 MB

func climbStairs(n int) int {
	//f(n)=f(n-1)+f(n-2)
	if n==1{
		return 1
	}else if n==2{
		return 2
	}
	fib:=make([]int,n+1)
	fib[1]=1
	fib[2]=2
	for i:=3;i<n+1;i++{
		fib[i]=fib[i-1]+fib[i-2]
	}
	return fib[n]

}
