package main

import (
	"os"
	"fmt"
	"io"
	"image"
	"image/jpeg"
)

// jpeg 命令从标准输入读取PNG图像
//并把它作为JPEG图像写到标准输出

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg:%v\n", err)
		os.Exit(1)
	}	
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
