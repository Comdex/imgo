package imgo

import(
	"encoding/base64"
	"io/ioutil"
)
//Img2Base64 produce a base64 string from a image file. 
func Img2Base64(filepath string)(encodeString string,err error){
	data,err:=ioutil.ReadFile(filepath)
	if err!=nil {
		return
	}
	encodeString=base64.StdEncoding.EncodeToString(data)
	return	
}

//Base64ToImg create a image file named dstFile from base64 encodeString.
func Base64ToImg(encodeString string,dstFile string)error{
	data,err:=base64.StdEncoding.DecodeString(encodeString)
	if err!=nil {
		return err
	}
	err2:=ioutil.WriteFile(dstFile,data,0666)
	if err2!=nil {
		return err2
	}
	return nil
}