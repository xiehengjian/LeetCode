
// @Title: 生成每种字符都是奇数个的字符串 (Generate a String With Characters That Have Odd Counts)
// @Author: 846188037@qq.com
// @Date: 2021-01-27 17:04:54
// @Runtime: 4 ms
// @Memory: 5.5 MB

func generateTheString(n int) string {
    res:=""
    if n%2==0{//偶数
    for i:=0;i<n-1;i++{
            res=res+"a"
        }
        res=res+"b"

    }else{//奇数
        for i:=0;i<n;i++{
            res=res+"a"
        }
    }
    return res
}
