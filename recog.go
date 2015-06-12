package imgo

import(
	"math"
)

//calculate Cosine Similarity of two images, input two file path
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

//binaryzation process of image matrix , threshold can use 127 to test
func Binaryzation(src [][][]uint8, threshold int)(imgMatrix [][][]uint8, err error) {
	imgMatrix,err = RGB2Gray(src)
	if err != nil {
		return
	}
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	for i:=0; i<height; i++ {
		for j:=0; j<width; j++ {
			var rgb int = int(imgMatrix[i][j][0])+int(imgMatrix[i][j][1])+int(imgMatrix[i][j][2])
			if rgb > threshold {
				rgb = 255
			}else{
				rgb = 0
			}
			imgMatrix[i][j][0]=uint8(rgb)
			imgMatrix[i][j][1]=uint8(rgb)
			imgMatrix[i][j][2]=uint8(rgb)
		}
	}
	
	return
}