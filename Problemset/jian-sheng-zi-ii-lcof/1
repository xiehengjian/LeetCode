
# @Title: 剪绳子 II (剪绳子 II LCOF)
# @Author: 846188037@qq.com
# @Date: 2021-04-18 11:17:23
# @Runtime: 956 ms
# @Memory: 15 MB

class Solution:
    def cuttingRope(self, n: int) -> int:
        f=[0 for i in range(n+1)]
        f[1]=1
        for i in range(2,n+1):
            for j in range(0,i):
                f[i]=max([f[i],f[j]*(i-j),j*(i-j)])

        return f[n]%1000000007

