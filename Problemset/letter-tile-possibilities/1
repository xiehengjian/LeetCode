
// @Title: 活字印刷 (Letter Tile Possibilities)
// @Author: 846188037@qq.com
// @Date: 2021-01-29 13:45:49
// @Runtime: 136 ms
// @Memory: 9.2 MB

func numTilePossibilities(tiles string) int {
    res:=make(map[string]int)
    temp:=""
    visited:=make([]bool,len(tiles))
    backtrack(tiles,temp,visited,&res)
    fmt.Println(res)
    return len(res)
}

func backtrack(tiles string,temp string,visited []bool, res *map[string]int){
    // if len(temp)==len(tiles){
    //     return
    // }
    for k,v :=range tiles{
        if visited[k]{
            continue
        }
        temp=temp+string(v)
        visited[k]=true
        (*res)[temp]++
        backtrack(tiles,temp,visited,res)
        temp=temp[:len(temp)-1]
        visited[k]=false
    }

}
