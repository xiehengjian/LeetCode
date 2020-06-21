#暴力循环，但是最终超时了
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        while True:
            num=nums[0]
            del nums[0]
            if num in nums:
                return num
#采用官方解法，还是超时           
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        num=[]
        for i in nums:
            if i in num:
                return i
            num.append(i)
            
#将暂存列表改为集合，高速通过
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        num=set()
        for i in nums:
            if i in num:
                return i
            num.add(i)

#省一个插入的步骤，还是超时
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            if nums[i] in nums[:i]:
                return nums[i]
            
 #故综上可知，集合的操作效率远远高于列表，所以同样的操作步骤，可以用集合就不要用列表       
