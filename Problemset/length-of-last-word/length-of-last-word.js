
// @Title: 最后一个单词的长度 (Length of Last Word)
// @Author: 846188037@qq.com
// @Date: 2020-07-09 20:17:35
// @Runtime: 60 ms
// @Memory: 31.8 MB

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    try{
    var list=s.split(" ")
    var i=1
    while (list[list.length-i].length==0)
    {
        i++
    }
    return list[list.length-i].length}
    catch(error){return 0}
};
