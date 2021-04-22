
// @Title: 括号生成 (Generate Parentheses)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 10:41:02
// @Runtime: 4 ms
// @Memory: 2.8 MB

func generateParenthesis(n int) []string {
    res:=make([]string,0)
    temp:=""
    Map:=make(map[string]int)
    Map["("]=0
    Map[")"]=0
    backtrack(n,Map,temp,&res)
    return res
}

func backtrack(n int,Map map[string]int,temp string,res *[]string){
    if len(temp)==2*n{
        *res=append(*res,temp)
        return
    }
    for k,v :=range Map{
        if Map[")"]>Map["("]{
            continue
        }
        if v>n-1{
            continue
        }
        temp=temp+k
        Map[k]++
        backtrack(n,Map,temp,res)
        Map[k]--
        temp=temp[:len(temp)-1]
    }
    }

