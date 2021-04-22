
// @Title: 矩阵中的路径 (矩阵中的路径  LCOF)
// @Author: 846188037@qq.com
// @Date: 2020-12-12 00:26:12
// @Runtime: 28 ms
// @Memory: 7.2 MB

func exist(board [][]byte, word string) bool {
    if len(board)==0 || len(board[0])==0{
        return false
    }
    visited:=make([][]bool,len(board))
    for i:=0;i<len(visited);i++{
        visited[i]=make([]bool,len(board[0]))
    }
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[0]);j++{
            if traverse(board,word,visited,i,j)==true{
                return true
            }

        }}
    return false
    
}


func traverse(board [][]byte,word string,visited [][]bool,i,j int)bool{
    if len(word)==0{
        return true
    }
    
    if i<0 || i>=len(board){
        return false
    }
    if j<0 || j>=len(board[0]){
        return false
    }
    if board[i][j]!=word[0]{
        return false
    }
    if visited[i][j]{
        return false
    }
    visited[i][j]=true
    return traverse(board,word[1:],visited,i+1,j) || traverse(board,word[1:],visited,i-1,j) || traverse(board,word[1:],visited,i,j+1) || traverse(board,word[1:],visited,i,j-1)

}
