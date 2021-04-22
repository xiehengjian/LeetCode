
// @Title: 单词规律 (Word Pattern)
// @Author: 846188037@qq.com
// @Date: 2021-03-23 09:14:13
// @Runtime: 0 ms
// @Memory: 2 MB

func wordPattern(pattern string, s string) bool {
    patternList:=strings.Split(pattern,"")
    strList:=strings.Split(s," ")
    return check(patternList,strList)&&check(strList,patternList)
}

func check(a []string,b []string) bool {
    Map:=make(map[string]string)
    if len(a)!=len(b){
        return false
    }
    for i:=0;i<len(a);i++{
        v,ok:=Map[a[i]]
        if !ok{
            Map[a[i]]=b[i]
        }else{
            if v!=b[i]{
                return false
            }
        }
    }
    return true
}
