
// @Title: 合并区间 (Merge Intervals)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 13:58:31
// @Runtime: 8 ms
// @Memory: 4.5 MB

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0]<intervals[j][0]
    })
    res:=make([][]int,0)
    pre:=intervals[0]
    for i:=1;i<len(intervals);i++{
        cur:=intervals[i]
        if pre[1]<cur[0]{
            res=append(res,pre)
            pre=cur
        }else{
            pre[1]=max(pre[1],cur[1])
        }
    }
    res=append(res,pre)
    return res

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
