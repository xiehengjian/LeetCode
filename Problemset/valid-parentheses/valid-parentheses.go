
// @Title: 有效的括号 (Valid Parentheses)
// @Author: 846188037@qq.com
// @Date: 2021-04-21 22:56:33
// @Runtime: 8 ms
// @Memory: 2.1 MB

func isValid(s string) bool {
    leftSign:=make([]byte,0)
    for i:=0;i<len(s);i++{
        if s[i]=='(' || s[i]=='[' || s[i]=='{'{
            leftSign=append(leftSign,s[i])
        }else{
            if len(leftSign)==0{
                return false
            }
            //fmt.Println(check(leftSign,s[i]))
            if !check(leftSign,s[i]){
                return false
            }
            leftSign=leftSign[:len(leftSign)-1]
        }
    }
    if len(leftSign)==0{
        return true
    }
    return false
}

func check(leftSign []byte,target byte)bool{
    n:=leftSign[len(leftSign)-1]
    fmt.Print(string(n))
    fmt.Println(string(target))
    switch target{
        case ')':
            if n=='('{
                return true
            }
        case ']':
            if n=='['{
                return true
            }
        case '}':
            if n=='{'{
                return true
            }
    }
    return false
}
