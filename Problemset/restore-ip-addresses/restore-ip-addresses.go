
// @Title: 复原 IP 地址 (Restore IP Addresses)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 20:53:59
// @Runtime: 4 ms
// @Memory: 2.9 MB

func restoreIpAddresses(s string) []string {
    if len(s)>12{
        return []string{}
    }
    res:=make([]string,0)
    temp:=""
    backtrack(s,len(s),temp,&res)
    return res
}

func backtrack(s string,n int,temp string,res *[]string){
    if len(temp)==n+4 && len(s)==0{
        
        *res=append(*res,temp[:len(temp)-1])
        return
    }
    for i:=1;i<4;i++{
        if len(s)<i {
            break
        }
        num,_:=strconv.Atoi(string(s[:i]))
        if num>255{
            continue
        }
        if i>1 && s[0]=='0'{
            continue
        }
        temp=temp+s[:i]+"."
        fmt.Println(temp)
        backtrack(s[i:],n,temp,res)
        temp=temp[:len(temp)-i-1]


    }


}
