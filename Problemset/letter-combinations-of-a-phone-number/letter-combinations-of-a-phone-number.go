
// @Title: 电话号码的字母组合 (Letter Combinations of a Phone Number)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 23:45:17
// @Runtime: 0 ms
// @Memory: 2 MB

func letterCombinations(digits string) []string {
    if len(digits)==0{
        return []string{}
    }
    res:=[]string{}
    temp:=""
    Map:=make(map[byte]string)
	Map['2']="abc"
	Map['3']="def"
	Map['4']="ghi"
	Map['5']="jkl"
	Map['6']="mno"
	Map['7']="pqrs"
	Map['8']="tuv"
	Map['9']="wxyz"

    backtrack(digits, Map,temp,&res)
    return res

}

func backtrack(digits string,Map map[byte]string, temp string,res *[]string){
    if len(digits)==0{
        *res=append(*res,temp)
        return 
    }

    for _,v :=range Map[digits[0]]{
        temp=temp+string(v)
       // fmt.Println(temp)
        backtrack(digits[1:],Map,temp,res)
        temp=temp[:len(temp)-1]
    }

}
