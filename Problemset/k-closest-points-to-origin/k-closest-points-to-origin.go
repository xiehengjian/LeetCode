
// @Title: 最接近原点的 K 个点 (K Closest Points to Origin)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 16:31:50
// @Runtime: 144 ms
// @Memory: 8.1 MB

func kClosest(points [][]int, K int) [][]int {
    sort.Slice(points,func(i,j int)bool{
        return points[i][0]*points[i][0]+points[i][1]*points[i][1]<points[j][0]*points[j][0]+points[j][1]*points[j][1]
    })
    return points[:K]
}
