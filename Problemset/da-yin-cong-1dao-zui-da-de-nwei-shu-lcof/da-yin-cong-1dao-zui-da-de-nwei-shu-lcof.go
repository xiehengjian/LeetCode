
// @Title: 打印从1到最大的n位数 (打印从1到最大的n位数 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 11:35:24
// @Runtime: 12 ms
// @Memory: 7 MB

func printNumbers(n int) []int {
    maxNum:=0
    for i:=0;i<n;i++{
        maxNum=maxNum*10+9
    }
    res:=make([]int,maxNum)
    for i:=0;i<maxNum;i++{
        res[i]=i+1
    }
    return res
    


}
