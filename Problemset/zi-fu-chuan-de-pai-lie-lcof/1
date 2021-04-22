
// @Title: 字符串的排列 (字符串的排列  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 15:14:06
// @Runtime: 128 ms
// @Memory: 9 MB

func permutation(s string) []string {
    res:=make(map[string]int,0)
    temp:=""
    visited:=make([]bool,len(s))
    backtrack(&res,temp,s,visited)
    ans:=make([]string,0)
    for k,_:=range res{
        ans=append(ans,k)
    }
    return ans


}

func backtrack(res *map[string]int,temp string,s string,visited []bool){
    if len(temp)==len(s){
        (*res)[temp]++
        return
    }
    for i:=0;i<len(s);i++{
        if visited[i]{
            continue
        }
        visited[i]=true
        temp+=string(s[i])
        backtrack(res,temp,s,visited)
        visited[i]=false
        temp=temp[:len(temp)-1]
    }

}
