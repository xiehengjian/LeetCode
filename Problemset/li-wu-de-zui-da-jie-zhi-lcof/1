
// @Title: 礼物的最大价值 (礼物的最大价值 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 15:50:53
// @Runtime: 8 ms
// @Memory: 4.3 MB

func maxValue(grid [][]int) int {
    if len(grid)==0 || len(grid[0])==0{
        return 0
    }
    f:=make([][]int,len(grid))
    for i:=0;i<len(grid);i++{
        f[i]=make([]int,len(grid[0]))
    }
    f[0][0]=grid[0][0]
    for i:=1;i<len(grid);i++{
        f[i][0]=f[i-1][0]+grid[i][0]
    }
    for i:=1;i<len(grid[0]);i++{
        f[0][i]=f[0][i-1]+grid[0][i]
    }
    for i:=1;i<len(grid);i++{
        for j:=1;j<len(grid[0]);j++{
            f[i][j]=max(f[i-1][j],f[i][j-1])+grid[i][j]
        }
    }
    return f[len(grid)-1][len(grid[0])-1]

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
