
// @Title: 整数反转 (Reverse Integer)
// @Author: 846188037@qq.com
// @Date: 2021-03-18 10:35:25
// @Runtime: 0 ms
// @Memory: 2.1 MB

func reverse(x int) int {
    flag:=x<0
    if flag{
        x=-x
    }
    res:=0
    for x>0{
        res=res*10+(x%10)
        x=x/10
        if res < -2147483648 || res > 2147483648{
            return 0
        }
    }
    if flag{
        return -res
    }
    return res

    
    
}
