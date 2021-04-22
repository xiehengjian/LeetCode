
// @Title: 重新排列数组 (Shuffle the Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-15 20:02:27
// @Runtime: 92 ms
// @Memory: 37.1 MB

/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number[]}
 */
var shuffle = function(nums, n) {
    var newNums= new Array()
    for(var i=0;i<n;i++){
        newNums.push(nums[i],nums[n+i])
    }
    return newNums


};
