
// @Title: 排序数组 (Sort an Array)
// @Author: 846188037@qq.com
// @Date: 2021-04-21 19:56:31
// @Runtime: 60 ms
// @Memory: 7.5 MB

func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}
func quickSort(a []int, l, r int) {
    if l >= r {
        return
    }
    a[r], a[(l+r)>>1] = a[(l+r)>>1], a[r]
    i := l - 1
    for j := l; j < r; j ++ {
        if a[j] < a[r] {
            i ++
            a[i], a[j] = a[j], a[i]
        }
    }
    i ++
    a[i], a[r] = a[r], a[i]
    quickSort(a, l, i-1 )
    quickSort(a, i+1, r)
}


