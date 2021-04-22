
# @Title: 一维数组的动态和 (Running Sum of 1d Array)
# @Author: 846188037@qq.com
# @Date: 2020-07-14 20:41:03
# @Runtime: 68 ms
# @Memory: 13.3 MB

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        runningSum=list()
        for i in range(len(nums)):
            runningSum.append(sum(nums[:i+1]))
        return runningSum
