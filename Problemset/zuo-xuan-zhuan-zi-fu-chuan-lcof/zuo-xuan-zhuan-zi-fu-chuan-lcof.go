
// @Title: 左旋转字符串 (左旋转字符串 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 17:32:44
// @Runtime: 0 ms
// @Memory: 3.1 MB

func reverseLeftWords(s string, n int) string {
    return s[n:]+s[:n]
}
