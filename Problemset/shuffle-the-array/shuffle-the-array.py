
# @Title: 重新排列数组 (Shuffle the Array)
# @Author: 846188037@qq.com
# @Date: 2020-07-15 19:56:40
# @Runtime: 36 ms
# @Memory: 13.5 MB

class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        newNums=list()
        for i in range(n):
            newNums.extend((nums[i],nums[n+i]))
        return newNums
