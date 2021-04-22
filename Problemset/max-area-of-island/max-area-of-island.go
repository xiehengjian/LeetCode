
// @Title: 岛屿的最大面积 (Max Area of Island)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 00:33:42
// @Runtime: 100 ms
// @Memory: 6.9 MB

func maxAreaOfIsland(grid [][]int) int {
    if len(grid)==0|| len(grid[0])==0{
        return 0
    }
    res:=0

    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j]==1{
                fmt.Println(res)
                res=max(traverse(grid,i,j),res)
            }
        }
    }
    return res


}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

func traverse(grid [][]int,i,j int)int{
    queue:=make([][]int,0)
    queue=append(queue,[]int{i,j})
    count:=0
    for len(queue)!=0{
        fmt.Println(queue)
        count++
        i,j:=queue[0][0],queue[0][1]
       
        //fmt.Println(queue)
        grid[i][j]=0
        queue=queue[1:]
        if i>0 && grid[i-1][j]==1{
            queue=append(queue,[]int{i-1,j})
            grid[i-1][j]=0
        }
        if i<len(grid)-1 && grid[i+1][j]==1{
            queue=append(queue,[]int{i+1,j})
            grid[i+1][j]=0
        }
        if j>0 && grid[i][j-1]==1{
            queue=append(queue,[]int{i,j-1})
            grid[i][j-1]=0
        }
        if j<len(grid[0])-1 && grid[i][j+1]==1{
            queue=append(queue,[]int{i,j+1})
            grid[i][j+1]=0
        }
    }
    return count

}
