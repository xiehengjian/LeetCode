
// @Title: 最长公共前缀 (Longest Common Prefix)
// @Author: 846188037@qq.com
// @Date: 2020-06-20 09:04:58
// @Runtime: 116 ms
// @Memory: 35.6 MB

/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    var a=""
    if(strs.length==0 || strs[0].length==0)
    return a;
    for(var k=0;k<strs[0].length;k++)
    {
        for(var i=0;i<strs.length;i++)
        {
        if (strs[i].charAt(k)!=strs[0].charAt(k))
        return a;
        }
       
        a=a+strs[0].charAt(k);
        

    }
    return a;
    

};
