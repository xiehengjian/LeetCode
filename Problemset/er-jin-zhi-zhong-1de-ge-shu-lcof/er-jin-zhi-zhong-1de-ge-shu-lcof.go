
// @Title: 二进制中1的个数 (二进制中1的个数 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 11:24:22
// @Runtime: 0 ms
// @Memory: 1.9 MB

func hammingWeight(num uint32) int {
    if num==0{
        return 0
    }
    res:=0
    for num>0{
        if num & 1==1{
            res++
        }
        num =num>>1
    }
    return res
}
