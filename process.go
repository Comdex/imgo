package imgo

import (
	"errors"
)

//input a image as src , return a image matrix by sunseteffect process
func SunsetEffect(src string)(imgMatrix [][][]uint8 , err error) {
	imgMatrix,err = Read(src)
	
	if err != nil {
		return 
	}
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][1] = uint8( float64(imgMatrix[i][j][1]) * 0.7 )
			imgMatrix[i][j][2] = uint8( float64(imgMatrix[i][j][2]) * 0.7 )
		}
	}
	
	return
}

// input a image as src , return a image matrix by negativefilmeffect process
func NegativeFilmEffect(src string)(imgMatrix [][][]uint8 , err error) {
	imgMatrix,err = Read(src)
	
	if err != nil {
		return 
	}
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = 255 - imgMatrix[i][j][0]
			imgMatrix[i][j][1] = 255 - imgMatrix[i][j][1]
			imgMatrix[i][j][2] = 255 - imgMatrix[i][j][2]
		}
	}
	
	return
}

func AdjustBrightness(src string , light float64)(imgMatrix [][][]uint8 , err error) {
	imgMatrix,err = Read(src)
	
	if err != nil {
		return 
	}
	
	if light <= 0{
		err = errors.New("value of light must be more than 0")
		return
	}
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = uint8(float64(imgMatrix[i][j][0])*light)
			imgMatrix[i][j][1] = uint8(float64(imgMatrix[i][j][1])*light)
			imgMatrix[i][j][2] = uint8(float64(imgMatrix[i][j][2])*light)
		}
	}
	
	return
}

func ImageFusion(src1 string , src2 string)(imgMatrix [][][]uint8 , err error) {
	imgMatrix1,err1 := Read(src1)
	
	if err1 != nil {
		err = err1
		return 
	}
	
	
	height:=len(imgMatrix1)
	width:=len(imgMatrix1[0])
	
	imgMatrix2,err2 := ResizeForMatrix(src2,width,height)
	
	if err2 != nil {
		err = err2
		return
	}
		
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix1[i][j][0] = uint8(float64(imgMatrix1[i][j][0])*0.5)+uint8(float64(imgMatrix2[i][j][0])*0.5)
			imgMatrix1[i][j][1] = uint8(float64(imgMatrix1[i][j][1])*0.5)+uint8(float64(imgMatrix2[i][j][1])*0.5)
			imgMatrix1[i][j][2] = uint8(float64(imgMatrix1[i][j][2])*0.5)+uint8(float64(imgMatrix1[i][j][2])*0.5)
		}
	}
	imgMatrix = imgMatrix1
	return	
}