
// @Title: 在既定时间做作业的学生人数 (Number of Students Doing Homework at a Given Time)
// @Author: 846188037@qq.com
// @Date: 2020-08-04 21:05:11
// @Runtime: 0 ms
// @Memory: 2.3 MB

func busyStudent(startTime []int, endTime []int, queryTime int) int {
    count := 0
    for i:=0;i<len(startTime);i++{
        if startTime[i]<=queryTime && queryTime<= endTime[i]{
            count++
        }
    }
    return count

}
