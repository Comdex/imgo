package imgo

import (
	"errors"
	"os"

	"image/jpeg"
	"image/png"

	"image"
	"image/color"
)

func GetImageHeight(img image.Image) int {
	return img.Bounds().Max.Y
}

func GetImageWidth(img image.Image) int {
	return img.Bounds().Max.X
}

// decode a image and retrun golang image interface
func DecodeImage(filePath string) (img image.Image, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	img, _, err = image.Decode(reader)

	return
}

//convert image to NRGBA
func convertToNRGBA(src image.Image) *image.NRGBA {
	srcBounds := src.Bounds()
	dstBounds := srcBounds.Sub(srcBounds.Min)

	dst := image.NewNRGBA(dstBounds)

	dstMinX := dstBounds.Min.X
	dstMinY := dstBounds.Min.Y

	srcMinX := srcBounds.Min.X
	srcMinY := srcBounds.Min.Y
	srcMaxX := srcBounds.Max.X
	srcMaxY := srcBounds.Max.Y

	switch src0 := src.(type) {

	case *image.NRGBA:
		rowSize := srcBounds.Dx() * 4
		numRows := srcBounds.Dy()

		i0 := dst.PixOffset(dstMinX, dstMinY)
		j0 := src0.PixOffset(srcMinX, srcMinY)

		di := dst.Stride
		dj := src0.Stride

		for row := 0; row < numRows; row++ {
			copy(dst.Pix[i0:i0+rowSize], src0.Pix[j0:j0+rowSize])
			i0 += di
			j0 += dj
		}

	case *image.NRGBA64:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				j := src0.PixOffset(x, y)

				dst.Pix[i+0] = src0.Pix[j+0]
				dst.Pix[i+1] = src0.Pix[j+2]
				dst.Pix[i+2] = src0.Pix[j+4]
				dst.Pix[i+3] = src0.Pix[j+6]

			}
		}

	case *image.RGBA:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				j := src0.PixOffset(x, y)
				a := src0.Pix[j+3]
				dst.Pix[i+3] = a

				switch a {
				case 0:
					dst.Pix[i+0] = 0
					dst.Pix[i+1] = 0
					dst.Pix[i+2] = 0
				case 0xff:
					dst.Pix[i+0] = src0.Pix[j+0]
					dst.Pix[i+1] = src0.Pix[j+1]
					dst.Pix[i+2] = src0.Pix[j+2]
				default:
					dst.Pix[i+0] = uint8(uint16(src0.Pix[j+0]) * 0xff / uint16(a))
					dst.Pix[i+1] = uint8(uint16(src0.Pix[j+1]) * 0xff / uint16(a))
					dst.Pix[i+2] = uint8(uint16(src0.Pix[j+2]) * 0xff / uint16(a))
				}
			}
		}

	case *image.RGBA64:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				j := src0.PixOffset(x, y)
				a := src0.Pix[j+6]
				dst.Pix[i+3] = a

				switch a {
				case 0:
					dst.Pix[i+0] = 0
					dst.Pix[i+1] = 0
					dst.Pix[i+2] = 0
				case 0xff:
					dst.Pix[i+0] = src0.Pix[j+0]
					dst.Pix[i+1] = src0.Pix[j+2]
					dst.Pix[i+2] = src0.Pix[j+4]
				default:
					dst.Pix[i+0] = uint8(uint16(src0.Pix[j+0]) * 0xff / uint16(a))
					dst.Pix[i+1] = uint8(uint16(src0.Pix[j+2]) * 0xff / uint16(a))
					dst.Pix[i+2] = uint8(uint16(src0.Pix[j+4]) * 0xff / uint16(a))
				}
			}
		}

	case *image.Gray:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				j := src0.PixOffset(x, y)
				c := src0.Pix[j]
				dst.Pix[i+0] = c
				dst.Pix[i+1] = c
				dst.Pix[i+2] = c
				dst.Pix[i+3] = 0xff

			}
		}

	case *image.Gray16:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				j := src0.PixOffset(x, y)
				c := src0.Pix[j]
				dst.Pix[i+0] = c
				dst.Pix[i+1] = c
				dst.Pix[i+2] = c
				dst.Pix[i+3] = 0xff

			}
		}

	case *image.YCbCr:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				yj := src0.YOffset(x, y)
				cj := src0.COffset(x, y)
				r, g, b := color.YCbCrToRGB(src0.Y[yj], src0.Cb[cj], src0.Cr[cj])

				dst.Pix[i+0] = r
				dst.Pix[i+1] = g
				dst.Pix[i+2] = b
				dst.Pix[i+3] = 0xff

			}
		}

	default:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {

				c := color.NRGBAModel.Convert(src.At(x, y)).(color.NRGBA)

				dst.Pix[i+0] = c.R
				dst.Pix[i+1] = c.G
				dst.Pix[i+2] = c.B
				dst.Pix[i+3] = c.A

			}
		}
	}

	return dst
}

// read a image return a image matrix by path or image
func Read(imgOrPath interface{}) (imgMatrix [][][]uint8, err error) {
	var img image.Image
	switch imgOrPath.(type) {
	case string:
		img, err = DecodeImage(imgOrPath.(string))
		if err != nil {
			return
		}
	case image.Image:
		img = imgOrPath.(image.Image)
	default:
		err = errors.New("Incoming parameter error!")
		return
	}

	bounds := img.Bounds()
	width := bounds.Max.X  //width
	height := bounds.Max.Y //height

	src := convertToNRGBA(img)
	imgMatrix = NewRGBAMatrix(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := src.At(j, i)
			r, g, b, a := c.RGBA()
			imgMatrix[i][j][0] = uint8(r)
			imgMatrix[i][j][1] = uint8(g)
			imgMatrix[i][j][2] = uint8(b)
			imgMatrix[i][j][3] = uint8(a)

		}
	}
	return
}

// read a image return a image matrix , if appear an error it will panic
func MustRead(filepath string) (imgMatrix [][][]uint8) {
	img, decodeErr := DecodeImage(filepath)
	if decodeErr != nil {
		panic(decodeErr)
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	src := convertToNRGBA(img)
	imgMatrix = NewRGBAMatrix(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := src.At(j, i)
			r, g, b, a := c.RGBA()
			imgMatrix[i][j][0] = uint8(r)
			imgMatrix[i][j][1] = uint8(g)
			imgMatrix[i][j][2] = uint8(b)
			imgMatrix[i][j][3] = uint8(a)

		}
	}
	return
}

// save a image matrix as a png , if unsuccessful it will return a error
func SaveAsPNG(filepath string, imgMatrix [][][]uint8) error {
	height := len(imgMatrix)
	width := len(imgMatrix[0])

	if height == 0 || width == 0 {
		return errors.New("The input of matrix is illegal!")
	}

	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{imgMatrix[i][j][0], imgMatrix[i][j][1], imgMatrix[i][j][2], imgMatrix[i][j][3]})
		}
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	png.Encode(outfile, nrgba)

	return nil
}

// save a image matrix as a jpeg,if unsuccessful it will return a error,quality must be 1 to 100
func SaveAsJPEG(filepath string, imgMatrix [][][]uint8, quality int) error {
	height := len(imgMatrix)
	width := len(imgMatrix[0])

	if height == 0 || width == 0 {
		return errors.New("The input of matrix is illegal!")
	}

	if quality < 1 {
		quality = 1
	} else if quality > 100 {
		quality = 100
	}

	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{imgMatrix[i][j][0], imgMatrix[i][j][1], imgMatrix[i][j][2], imgMatrix[i][j][3]})
		}
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	jpeg.Encode(outfile, nrgba, &jpeg.Options{Quality: quality})

	return nil
}
