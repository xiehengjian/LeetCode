
// @Title: 最后一个单词的长度 (Length of Last Word)
// @Author: 846188037@qq.com
// @Date: 2021-04-02 08:19:14
// @Runtime: 0 ms
// @Memory: 2.3 MB

func lengthOfLastWord(s string) int {
    list:=strings.Split(s," ")
    fmt.Println(list)
    for i:=len(list)-1;i>=0;i--{
        fmt.Print(i)
        fmt.Println(list[i])
        if list[i]!=""{
            return len(list[i])
        }
    }
    return 0

}
