#暴力循环，但是最终超时了
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        while True:
            num=nums[0]
            del nums[0]
            if num in nums:
                return num
