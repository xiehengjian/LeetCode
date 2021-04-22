
// @Title: 二维数组中的查找 (二维数组中的查找 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 00:32:05
// @Runtime: 32 ms
// @Memory: 6.7 MB

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix)==0||len(matrix[0])==0{
        return false
    }
    return helper(matrix,target,0,len(matrix[0])-1)
    
}

func helper(matrix [][]int,target int, m,n int)bool{
    if m>len(matrix)-1 || n <0{
        return false
    }
    if matrix[m][n]>target{
        return helper(matrix,target,m,n-1)
    }else if matrix[m][n]<target{
        return helper(matrix,target,m+1,n)
    }
    return true

}



