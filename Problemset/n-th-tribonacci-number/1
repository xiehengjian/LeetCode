
// @Title: 第 N 个泰波那契数 (N-th Tribonacci Number)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 13:24:33
// @Runtime: 0 ms
// @Memory: 1.9 MB

func tribonacci(n int) int {
	if n==0{
		return 0
	}else if n==1||n==2{
		return 1
	}
	fib:=make([]int,n+1)
	fib[0]=0
	fib[1]=1
	fib[2]=1
	for i:=3;i<n+1;i++{
		fib[i]=fib[i-1]+fib[i-2]+fib[i-3]
		
	}
	
	return fib[n]

}
