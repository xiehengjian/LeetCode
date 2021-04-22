
// @Title: 一维数组的动态和 (Running Sum of 1d Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-14 21:15:26
// @Runtime: 100 ms
// @Memory: 34.9 MB

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function(nums) {
    var sum=0;
    var runningSum = new Array;
    for(var i=0;i<nums.length;i++){
        for(var j=0;j<i+1;j++){
            sum+=nums[j];
        }
        runningSum.push(sum);
        sum=0;
    }
    return runningSum;

};
