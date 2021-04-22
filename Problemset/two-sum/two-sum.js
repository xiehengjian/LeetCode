
// @Title: 两数之和 (Two Sum)
// @Author: 846188037@qq.com
// @Date: 2020-05-22 11:53:37
// @Runtime: 140 ms
// @Memory: 34.7 MB

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    for(var i=0;i<nums.length-1;i++)
    {
        for(var j=i+1;j<nums.length;j++)
        {
            if(nums[i]+nums[j]==target)
            {
                var ans=[i,j]
                return ans
            }
        }
    }

};
