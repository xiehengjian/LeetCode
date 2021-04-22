
# @Title: 最后一个单词的长度 (Length of Last Word)
# @Author: 846188037@qq.com
# @Date: 2020-07-09 20:04:53
# @Runtime: 32 ms
# @Memory: 13.3 MB

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        try:
            return len(s.split()[-1])
        except:
            return 0
