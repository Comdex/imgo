package imgo

import (
	"errors"
)

//input a image matrix as src , return a image matrix by sunseteffect process
func SunsetEffect(src [][][]uint8) [][][]uint8 {
	
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][1] = uint8( float64(imgMatrix[i][j][1]) * 0.7 )
			imgMatrix[i][j][2] = uint8( float64(imgMatrix[i][j][2]) * 0.7 )
		}
	}
	
	return imgMatrix
}

// input a image as src , return a image matrix by negativefilmeffect process
func NegativeFilmEffect(src [][][]uint8) [][][]uint8 {
	
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = 255 - imgMatrix[i][j][0]
			imgMatrix[i][j][1] = 255 - imgMatrix[i][j][1]
			imgMatrix[i][j][2] = 255 - imgMatrix[i][j][2]
		}
	}
	
	return imgMatrix
}

func Rotate(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(width,height)
	
	for i:=0;i<width;i++{
		for j:=0;j<height;j++{
			imgMatrix[i][j] = src[j][i]
		}
	}
	
	return imgMatrix
}

func AdjustBrightness(src [][][]uint8 , light float64)(imgMatrix [][][]uint8 , err error) {
	
	if light <= 0{
		err = errors.New("value of light must be more than 0")
		return
	}
	
	height := len(src)
	width := len(src[0])
	imgMatrix = NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = uint8(float64(imgMatrix[i][j][0])*light)
			imgMatrix[i][j][1] = uint8(float64(imgMatrix[i][j][1])*light)
			imgMatrix[i][j][2] = uint8(float64(imgMatrix[i][j][2])*light)
		}
	}
	
	return
}

// fuse two images(filepath) and the size of new image is as src1
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

func VerticalMirror(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	
	newwidth:=width*2
	imgMatrix:=NewRGBAMatrix(height,newwidth)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		imgMatrix[i][j] = src[i][j]
		}
	}
	
	
	for i:=0;i<height;i++{
		for j:=width;j<newwidth;j++{
			imgMatrix[i][j] = imgMatrix[i][newwidth-j-1]
		}
	}
	
	return imgMatrix
}

func HorizontalMirror(src [][][]uint8) [][][]uint8 {
	height:=len(src)
	width:=len(src[0])
	
	newheight:=height*2
	imgMatrix:=NewRGBAMatrix(newheight,width)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		imgMatrix[i][j] = src[i][j]
		}
	}
	
	
	for i:=height;i<newheight;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j] = imgMatrix[newheight-i-1][j]
		}
	}
	
	return imgMatrix
}


func VerticalMirrorPart(src [][][]uint8) [][][]uint8 {
	
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	mirror_w:=width/2
	
	for i:=0;i<height;i++{
		for j:=0;j<mirror_w;j++{
			imgMatrix[i][j] = imgMatrix[i][width-j-1]
		}
	}
	
	return imgMatrix
}

//make a mirror of src 
func HorizontalMirrorPart(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	mirror_h:=height/2
	
	for i:=0;i<mirror_h;i++{
		for j:=0;j<width;j++{
		imgMatrix[height-i-1][j] = imgMatrix[i][j]
		}
	}
	
	return imgMatrix
}


func RGB2Gray(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	imgMatrix := NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		avg:=(imgMatrix[i][j][0]+imgMatrix[i][j][1]+imgMatrix[i][j][3])/3
		imgMatrix[i][j][0] = avg
		imgMatrix[i][j][1] = avg
		imgMatrix[i][j][2] = avg
		}
	}
	return imgMatrix
}

// set the opacity of image matrix , opacity must be 0.0 to 1.0
func SetOpacity(src [][][]uint8, opacity float64)(imgMatrix [][][]uint8 , err error){
	height := len(src)
	width := len(src[0])
	imgMatrix = NewRGBAMatrix(height,width)
	copy(imgMatrix,src)
	
	if opacity < 0.0 || opacity > 1.0 {
		err = errors.New("the opacity is illegal!")
	}
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		imgMatrix[i][j][3] = uint8(float64(imgMatrix[i][j][3])*opacity)
		}
	}
	return
}