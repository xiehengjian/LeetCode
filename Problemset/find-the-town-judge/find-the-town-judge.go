
// @Title: 找到小镇的法官 (Find the Town Judge)
// @Author: 846188037@qq.com
// @Date: 2021-02-16 10:07:01
// @Runtime: 112 ms
// @Memory: 7.3 MB

func findJudge(N int, trust [][]int) int {
    in:=make([]int,N+1)
    out:=make([]int,N+1)
    for i:=0;i<len(trust);i++{
        out[trust[i][0]]++
        in[trust[i][1]]++
    }
    for i:=1;i<N+1;i++{
        if in[i]==N-1&&out[i]==0{
            return i
        }
    }
    return -1

}
