
// @Title: 单词拆分 (Word Break)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 22:37:40
// @Runtime: 0 ms
// @Memory: 2.4 MB

func wordBreak(s string, wordDict []string) bool {
	//f[i]前i个字符能否拆分成功
	//f[i]=f[j]&& f[j+1:i] in wordDict
	if len(s)==0{
		return true
	}
	n:=len(s)
	f:=make([]bool,n+1)
	Map:=make(map[string]int)
	for _,v:=range wordDict{
		Map[v]++
	}
	f[0]=true
	for i:=1;i<n+1;i++{
		for j:=0;j<i;j++{
			if _,ok:=Map[s[j:i]];f[j]&&ok{
				f[i]=true
			}
		}
	}
	fmt.Println(f)
	return f[n]

}
