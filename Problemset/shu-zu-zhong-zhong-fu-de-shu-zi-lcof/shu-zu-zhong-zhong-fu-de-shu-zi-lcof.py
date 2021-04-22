
# @Title: 数组中重复的数字 (数组中重复的数字 LCOF)
# @Author: 846188037@qq.com
# @Date: 2020-06-21 09:35:58
# @Runtime: 48 ms
# @Memory: 22.6 MB

class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        num=set()
        for i in nums:
            if i in num:
                return i
            num.add(i)


                

