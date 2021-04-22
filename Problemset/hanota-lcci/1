
// @Title: 汉诺塔问题 (Hanota LCCI)
// @Author: 846188037@qq.com
// @Date: 2020-11-21 14:05:11
// @Runtime: 36 ms
// @Memory: 2 MB

func hanota(A []int, B []int, C []int) []int {
	full:=len(A)
	step:=[2]int{0,0}
	for len(C)!=full{
		if (step[0]!=2||step[1]!=1)&&len(A)!=0&&(len(B)==0||A[len(A)-1]<B[len(B)-1]){
			step=[2]int{1,2}
			B=append(B,A[len(A)-1])
			A=A[:len(A)-1]
		}else if (step[0]!=1||step[1]!=2)&&len(B)!=0&&(len(A)==0||B[len(B)-1]<A[len(A)-1]){
			step=[2]int{2,1}
			A=append(A,B[len(B)-1])
			B=B[:len(B)-1]
		}else if( step[0]!=3||step[1]!=2)&&len(B)!=0&&(len(C)==0||B[len(B)-1]<C[len(C)-1]){
			step=[2]int{2,3}
			C=append(C,B[len(B)-1])
			B=B[:len(B)-1]
		}else if (step[0]!=2||step[1]!=3)&&len(C)!=0&&(len(B)==0||C[len(C)-1]<B[len(B)-1]){
			step=[2]int{3,2}
			B=append(B,C[len(C)-1])
			C=C[:len(C)-1]
		}
	}
	return C

}
