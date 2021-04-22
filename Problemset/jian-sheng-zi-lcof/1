
// @Title: 剪绳子 (剪绳子  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 10:50:20
// @Runtime: 0 ms
// @Memory: 2 MB

func cuttingRope(n int) int {
    //f[i]长度为i时的最大乘积
    //f[i]=max(f[j]*(i-j))
    f:=make([]int,n+1)
    f[1]=1
    for i:=2;i<n+1;i++{
        for j:=0;j<i;j++{
            f[i]=max(f[i],f[j]*(i-j))
            f[i]=max(f[i],j*(i-j))
            
        }
    }
    fmt.Println(f)
    return f[n]


}

func max(a,b int)int{
    if a>b{
        return a 
    }
    return b
}
