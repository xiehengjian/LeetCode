
// @Title: x 的平方根 (Sqrt(x))
// @Author: 846188037@qq.com
// @Date: 2021-04-20 00:00:15
// @Runtime: 4 ms
// @Memory: 2.2 MB

func mySqrt(x int) int {
    if x<1{
        return 0
    }
    if x==1{
        return 1
    }
    left,right:=0,x
    for left<right{
        mid:=(left+right)/2
        if mid*mid==x{
            return mid
        }else if mid*mid>x{
            if (mid-1)*(mid-1)<=x{
                return mid-1
            }
            right=mid
        }else if mid*mid<x{
            if (mid+1)*(mid+1)>x{
                return mid
            }
            left=mid
        }
    }
    return 0
}
