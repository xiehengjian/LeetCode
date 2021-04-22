
// @Title: 接雨水 (Trapping Rain Water)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 20:05:22
// @Runtime: 4 ms
// @Memory: 3.1 MB

func trap(height []int) int {
    if len(height)==0{
        return 0
    }
    left:=make([]int,len(height))
    left[0]=height[0]
    for i:=1;i<len(height);i++{
        left[i]=max(left[i-1],height[i])
    }
    right:=make([]int,len(height))
    right[len(height)-1]=height[len(height)-1]
    for i:=len(height)-2;i>-1;i--{
        right[i]=max(right[i+1],height[i])
    }
    ans:=0
    for i:=0;i<len(height);i++{
        ans+=min(left[i],right[i])-height[i]
    }
    return ans

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}
