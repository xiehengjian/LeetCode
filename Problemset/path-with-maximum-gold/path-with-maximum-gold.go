
// @Title: 黄金矿工 (Path with Maximum Gold)
// @Author: 846188037@qq.com
// @Date: 2021-01-29 12:58:36
// @Runtime: 156 ms
// @Memory: 6.9 MB

func getMaximumGold(grid [][]int) int {
    res:=0
    temp:=0
    visited:=make([][]bool,len(grid))
    for i:=0;i<len(visited);i++{
        visited[i]=make([]bool,len(grid[0]))
    }
    list:=[][]int{}
    count:=0
    backtrack(grid,count,visited,list,temp,&res)
    return res
}

func backtrack(grid [][]int,count int,visited [][]bool, list [][]int,temp int,res *int){
       if temp>(*res){
            *res=temp
        }
   // if count>=len(grid)*len(grid[0]){
        // if temp>(*res){
        //     *res=temp
        // }
        //return 
   // }
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            count++
            if grid[i][j]==0{
                continue
            }
            if visited[i][j]{
                continue
            }


             if len(list)!=0 &&!((list[len(list)-1][0]==i && (list[len(list)-1][1]==j+1 || list[len(list)-1][1]==j-1)) || (list[len(list)-1][1]==j && (list[len(list)-1][0]==i+1 || list[len(list)-1][0]==i-1))) {
                 continue
             }
            list=append(list,[]int{i,j})
            visited[i][j]=true
            temp=temp+grid[i][j]
          
            backtrack(grid,count,visited,list,temp,res)
            list=list[:len(list)-1]
            visited[i][j]=false
            temp=temp-grid[i][j]
            
        }
    }

}
