
# @Title: 好数对的数目 (Number of Good Pairs)
# @Author: 846188037@qq.com
# @Date: 2020-07-12 21:06:58
# @Runtime: 48 ms
# @Memory: 13.5 MB

class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        count=0
        for i in range(len(nums)-1):
            for j in range(i+1,len(nums)):
                if nums[i]==nums[j]:
                    count+=1
        return count
