
// @Title: 整数的各位积和之差 (Subtract the Product and Sum of Digits of an Integer)
// @Author: 846188037@qq.com
// @Date: 2020-08-03 15:37:16
// @Runtime: 0 ms
// @Memory: 1.9 MB

func subtractProductAndSum(n int) int {
    s := strconv.Itoa(n)
    sum := 0
    pro := 1
    for _,v := range s{
        a,_ := strconv.Atoi(string(v))
        sum += a
        b,_ := strconv.Atoi(string(v))
        pro *= b
    }
    return pro - sum

}
