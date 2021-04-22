
// @Title: 岛屿数量 (Number of Islands)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 15:30:03
// @Runtime: 0 ms
// @Memory: 2.7 MB

func numIslands(grid [][]byte) int {
    count:=0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j]=='1'{
                count++
                traverse(&grid,i,j)
            }
        }
    }
    return count

}

func traverse(grid *[][]byte,i,j int){
    (*grid)[i][j]='0'
    if i-1>=0 && (*grid)[i-1][j]=='1'{
        traverse(grid,i-1,j)
    }
    if i+1<len((*grid)) && (*grid)[i+1][j]=='1'{
        traverse(grid,i+1,j)
    }
    if j-1>=0 && (*grid)[i][j-1]=='1'{
        traverse(grid,i,j-1)
    }
    if j+1<len((*grid)[0]) && (*grid)[i][j+1]=='1'{
        traverse(grid,i,j+1)
    }

    
    
    
    

}
