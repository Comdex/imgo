package imgo

import(
	"math"
)

func CosineSimilarity(src1 string, src2 string)(cossimi float64,err error){
	matrix1,err1:=ResizeForMatrix(src1,80,60)
	if err1 != nil {
		err = err1
		return
	}
	
	matrix2,err2:=ResizeForMatrix(src2,80,60)
	if err2 != nil {
		err = err2
		return
	}
	
	myx:=Matrix2Vector(matrix1)
	myy:=Matrix2Vector(matrix2)
	cos1:=Dot(myx,myy)
	cos21:=math.Sqrt(Dot(myx,myx))
	cos22:=math.Sqrt(Dot(myy,myy))
	
	cossimi = cos1/(cos21*cos22)
	return
}