# imgo
golang图像处理工具库(golang image process lib)

[![GoDoc](http://godoc.org/github.com/Comdex/imgo?status.svg)](http://godoc.org/github.com/Comdex/imgo)

### 安装

```shell
go get github.com/Comdex/imgo
```

### 示例

```go
package main

import(
	"github.com/Comdex/imgo"
)

func main(){
    //如果读取出错会panic,返回图像矩阵
    //img[height][width][4],height为图像高度，width为图像宽度
    //img[height][width][4]为第height行第width列上像素点的RGBA数值数组，值范围为0-255
    img:=imgo.MustRead("test.jpg")
	
	//对原图像矩阵进行日落效果处理
	img2,err:=imgo.SunsetEffect(img)
	if err!=nil {
		panic(err)
	}
	
	//保存为jpeg
	err2:=imgo.SaveAsJPEG("new.jpg",img2)
	if err2!=nil {
		panic(err)
	}
}
```

### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。

