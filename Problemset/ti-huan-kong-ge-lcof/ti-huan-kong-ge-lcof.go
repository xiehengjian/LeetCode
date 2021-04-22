
// @Title: 替换空格 (替换空格 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 00:39:30
// @Runtime: 0 ms
// @Memory: 3.4 MB

func replaceSpace(s string) string {
    res:=""
    for i:=0;i<len(s);i++{
        if s[i]==' '{
            res+="%20"
        }else{
            res+=string(s[i])
        }
    }
    return res
}
