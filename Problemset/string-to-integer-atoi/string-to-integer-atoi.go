
// @Title: 字符串转换整数 (atoi) (String to Integer (atoi))
// @Author: 846188037@qq.com
// @Date: 2021-04-19 10:56:18
// @Runtime: 0 ms
// @Memory: 2.2 MB

func myAtoi(s string) int {
    if len(s)==0{
        return 0
    }
    for i:=0;i<len(s);i++{
        if s[i]!=' '{
            s=s[i:]
            break
        }
    }
    flag:=1
    if s[0]=='-' || s[0]=='+'{
        if s[0]=='-'{
            flag=-1
        }
        s=s[1:]
    }
    res:=0

    for i:=0;i<len(s);i++{
        if s[i]>='0' && s[i]<='9'{
            temp:=0-int('0')+int(s[i])
            res=res*10+temp
             if float64(res*flag)<(-math.Pow(2,31)){
                return int(-math.Pow(2,31))
                } 
                if float64(res*flag)>math.Pow(2,31)-1{
                    return int(math.Pow(2,31)-1)
                }
        }else{
            break
        }
    }
    
    return res*flag
}
