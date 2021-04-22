
// @Title: 宝石与石头 (Jewels and Stones)
// @Author: 846188037@qq.com
// @Date: 2021-01-29 17:15:38
// @Runtime: 0 ms
// @Memory: 2 MB

func numJewelsInStones(J string, S string) int {
    count := 0
    for _,v := range S{
        if strings.Contains(J,string(v)){
            count++
        }
    }
    return count

}
