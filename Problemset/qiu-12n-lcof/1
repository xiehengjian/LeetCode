
// @Title: 求1+2+…+n (求1+2+…+n LCOF)
// @Author: 846188037@qq.com
// @Date: 2020-08-01 15:19:12
// @Runtime: 0 ms
// @Memory: 2.5 MB

func sumNums(n int) int {
    ans := 0
    var sum func(int) bool
    sum = func(n int) bool{
        ans+=n
        return n>0&&sum(n-1)
    }
    sum(n)
    return ans
}
