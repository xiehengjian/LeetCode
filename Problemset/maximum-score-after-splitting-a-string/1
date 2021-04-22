
// @Title: 分割字符串的最大得分 (Maximum Score After Splitting a String)
// @Author: 846188037@qq.com
// @Date: 2021-01-27 16:53:41
// @Runtime: 0 ms
// @Memory: 1.9 MB

func maxScore(s string) int {
    count_of_one:=0
    for i:=0;i<len(s);i++{
        if s[i]=='1'{
            count_of_one++
        }
    }
    left:=0
    right:=count_of_one
    score:=0
    for i:=0;i<len(s)-1;i++{
        if s[i]=='0'{
            left++
        }else{
            right--
        }
        if left+right>score{
            score=left+right
        }

    }
    return score

}
