
// @Title: 螺旋矩阵 (Spiral Matrix)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 14:43:54
// @Runtime: 0 ms
// @Memory: 2 MB

func spiralOrder(matrix [][]int) []int {
    
    res:=make([]int,0)
    traverse(matrix,&res)
    return res

}


func traverse(matrix [][]int,res *[]int){
    fmt.Println(matrix)
    if len(matrix)==0 || len(matrix[0])==0{
        return
    }//s双数，最后为空，返回
    if len(matrix)==1 || len(matrix[0])==1{//剩一行，直接添加后返回
        for _,row:=range matrix{
            for _,e:=range row{
                *res=append(*res,e)
            }
        }
    return
    }

    

    for i:=0;i<len(matrix[0])-1;i++{
        *res=append(*res,matrix[0][i])
    }
    for i:=0;i<len(matrix)-1;i++{
        *res=append(*res,matrix[i][len(matrix[0])-1])
    }
    for i:=len(matrix[0])-1;i>0;i--{
        *res=append(*res,matrix[len(matrix)-1][i])
    }
    for i:=len(matrix)-1;i>0;i--{
        *res=append(*res,matrix[i][0])
    }
    
    matrix=matrix[1:]
    matrix=matrix[:len(matrix)-1]
    for i:=0;i<len(matrix);i++{
        matrix[i]=matrix[i][1:]
        matrix[i]=matrix[i][:len(matrix[i])-1]
    }
    traverse(matrix,res)
}
