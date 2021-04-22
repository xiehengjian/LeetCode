
// @Title: 排列序列 (Permutation Sequence)
// @Author: 846188037@qq.com
// @Date: 2021-04-08 21:44:46
// @Runtime: 0 ms
// @Memory: 2 MB

func getPermutation(n int, k int) string {
    factorial := make([]int, n)
    factorial[0] = 1
    for i := 1; i < n; i++ {
        factorial[i] = factorial[i - 1] * i
    }
    k--

    ans := ""
    valid := make([]int, n + 1)
    for i := 0; i < len(valid); i++ {
        valid[i] = 1
    }
    for i := 1; i <= n; i++ {
        order := k / factorial[n - i] + 1
        for j := 1; j <= n; j++ {
            order -= valid[j]
            if order == 0 {
                ans += strconv.Itoa(j)
                valid[j] = 0
                break
            }
        }
        k %= factorial[n - i]
    }
    return ans
}


