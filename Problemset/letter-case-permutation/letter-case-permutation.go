
// @Title: 字母大小写全排列 (Letter Case Permutation)
// @Author: 846188037@qq.com
// @Date: 2021-04-09 15:20:08
// @Runtime: 12 ms
// @Memory: 6.4 MB

func letterCasePermutation(S string) []string {
	res:=make([]string,0)
	temp:=""
	visited:=make([]bool,len(S))
	backtrack(&res,temp,visited,S)
	return res

}
func backtrack(res *[]string,temp string,visited []bool,S string){
	if len(S)==0{
		*res=append(*res,temp)
		return
	}

	sub:=S[0]
	//S=S[1:]
	if sub>='a' && sub<='z'{
		temp=temp+string(sub+'A'-'a')
		backtrack(res,temp,visited,S[1:])
		temp=temp[:len(temp)-1]
	}else if sub>='A' && sub<='Z'{
		temp=temp+string(sub+'a'-'A')
		backtrack(res,temp,visited,S[1:])
		temp=temp[:len(temp)-1]
	}
	temp=temp+string(sub)
	backtrack(res,temp,visited,S[1:])
	temp=temp[:len(temp)-1]


}
