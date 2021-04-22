
// @Title: 分割回文串 (Palindrome Partitioning)
// @Author: 846188037@qq.com
// @Date: 2021-01-29 14:42:02
// @Runtime: 340 ms
// @Memory: 22.1 MB

func partition(s string) [][]string {
    res:=[][]string{}
    temp:=[]string{}
    backtrack(s,temp,&res)
    return res

}

func backtrack(s string,temp []string,res *[][]string){
    if len(s)==0{
        ans := make([]string, len(temp))
        copy(ans, temp)
        *res=append(*res,ans)
        return
    }
    for i:=1;i<len(s)+1;i++{
        sub:=s[:i]
        if !isPalindrome(sub){
            continue
        }
        temp=append(temp,sub)
        backtrack(s[i:],temp,res)
        temp=temp[:len(temp)-1]
    }

}

func isPalindrome(s string) bool {
    left:=0
    right:=len(s)-1
    for left<right{
        if !strings.EqualFold(string(s[left]),string(s[right])){
            return false
        }
        left++
        right--
    }
    return true
}

