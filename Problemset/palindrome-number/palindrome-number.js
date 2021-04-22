
// @Title: 回文数 (Palindrome Number)
// @Author: 846188037@qq.com
// @Date: 2020-06-08 21:15:46
// @Runtime: 224 ms
// @Memory: 43.4 MB

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    
    x=String(x)
    if(x.charAt(0)=="-")
    {
        return false
    }
    var flag=0;
    for(var i=Math.round(x.length/2)-1;i>=0;i--)
    {
        if(x.charAt(i)!==x.charAt(x.length-1-i))
        {
            flag=1;
            return false
        }
    }
    return true
  

};
