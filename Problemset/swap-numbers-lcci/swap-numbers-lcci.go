
// @Title: 交换数字 (Swap Numbers LCCI)
// @Author: 846188037@qq.com
// @Date: 2020-08-03 15:56:28
// @Runtime: 0 ms
// @Memory: 1.9 MB

func swapNumbers(numbers []int) []int {
    numbers[0],numbers[1] = numbers[1],numbers[0] 
    return numbers

}
