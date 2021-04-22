
// @Title: 不同路径 II (Unique Paths II)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 20:11:49
// @Runtime: 4 ms
// @Memory: 2.7 MB

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	
	if len(obstacleGrid)==0 || len(obstacleGrid[0])==0{
		return 0
	}
	if obstacleGrid[0][0]==1{
		return 0
	}
	m:=len(obstacleGrid)
	n:=len(obstacleGrid[0])
	f:=make([][]int,m)
	for i:=0;i<m;i++{
		f[i]=make([]int,n)	
	}
	f[0][0]=1

	for i:=1;i<m;i++{
		if obstacleGrid[i][0]==1{
			f[i][0]=0
		}else{
			f[i][0]=f[i-1][0]
		}
		
	}
	for i:=1;i<n;i++{
		if obstacleGrid[0][i]==1{
			f[0][i]=0
		}else{
			f[0][i]=f[0][i-1]
		}
		
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if obstacleGrid[i][j]==1{
				f[i][j]=0
			}else{
				f[i][j]=f[i-1][j]+f[i][j-1]
			}
			
		}
	}
	fmt.Println(f)
	return f[m-1][n-1]

}
