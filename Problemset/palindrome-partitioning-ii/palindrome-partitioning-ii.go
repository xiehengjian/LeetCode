
// @Title: 分割回文串 II (Palindrome Partitioning II)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 21:13:28
// @Runtime: 152 ms
// @Memory: 2.1 MB

func minCut(s string) int {
	if len(s)==0||len(s)==1{
		return 0
	}
	n:=len(s)
	f:=make([]int,n+1)
	f[0]=-1
	f[1]=0
	for i:=1;i<n+1;i++{
		f[i]=i-1
		for j:=0;j<i;j++{
			if isPalindrome(s,j,i-1){
				f[i]=min(f[i],f[j]+1)
			}
		}
	}
	return f[n]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
