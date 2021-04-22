
// @Title: 验证回文串 (Valid Palindrome)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 15:41:28
// @Runtime: 0 ms
// @Memory: 2.6 MB

func isPalindrome(s string) bool {
    left:=0
    right:=len(s)-1
    for left<right{
        for !isLetterOrNumber(s[left])&&left<right {
            left++
        }
        for !isLetterOrNumber(s[right])&&left<right{
            right--
        }
        if !strings.EqualFold(string(s[left]),string(s[right])){
            return false
        }
        left++
        right--

    }
    return true
}

func isLetterOrNumber(char byte)bool{
    return (char>='0'&&char<='9')||(char>='a'&&char<='z')||(char>='A'&&char<='Z')    
    
}
