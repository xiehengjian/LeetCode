
// @Title: 旋转数组的最小数字 (旋转数组的最小数字  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 09:25:12
// @Runtime: 4 ms
// @Memory: 3.1 MB

func minArray(numbers []int) int {
    if len(numbers)==0{
        return -1
    }
    if len(numbers)==1{
        return numbers[0]
    }
    temp:=numbers[0]
    for i:=1;i<len(numbers);i++{
        if numbers[i]<temp{
            return numbers[i]
        }
        temp=numbers[i]
    }
    return numbers[0]

}
