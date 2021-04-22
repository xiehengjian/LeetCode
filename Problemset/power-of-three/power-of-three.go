
// @Title: 3的幂 (Power of Three)
// @Author: 846188037@qq.com
// @Date: 2021-03-29 22:48:25
// @Runtime: 24 ms
// @Memory: 6.2 MB

func isPowerOfThree(n int) bool {
    if n<1{
        return false
    }
    for n%3==0{
        n=n/3
    }
    return n==1
}
