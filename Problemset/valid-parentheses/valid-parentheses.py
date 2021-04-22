
# @Title: 有效的括号 (Valid Parentheses)
# @Author: 846188037@qq.com
# @Date: 2020-07-11 20:43:20
# @Runtime: 52 ms
# @Memory: 13.4 MB

class Solution:
    def isValid(self, s: str) -> bool:
        try:
            s=list(s)
            s1=list()
            while len(s)!=0:
                if s[0]=="(" or s[0]=='{' or s[0]=="[":
                    s1.append(s.pop(0))
                if s[0] == ")" :
                    if s1.pop() =="(":
                        s.pop(0)
                        continue
                    else:
                        return False
                elif s[0] == "]":
                    if s1.pop() =="[":
                        s.pop(0)
                        continue
                    else:
                        return False
                elif s[0] == "}":
                    if s1.pop() =="{":
                        s.pop(0)
                        continue
                    else:
                        return False
                
            if len(s1)!=0:
                return False
            else:
                return True
        except:
            return False
                
