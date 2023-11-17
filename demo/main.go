package main

import (
	"fmt"
	"os"

	ppocr "github.com/skye-z/ppocr-go"
)

func main() {
	rootPath, _ := os.Getwd()
	fmt.Println("创建引擎")
	// 创建引擎
	ocr := ppocr.Engine{
		Config: ppocr.OCRConfig{
			UseGPU:         false,
			DBDetectorPath: rootPath + "/model/det-v4",
			ClassifierPath: rootPath + "/model/cls-v2",
			RecognizerPath: rootPath + "/model/rec-v4",
			KeysPath:       rootPath + "/model/keys.txt",
		},
	}
	fmt.Println("加载模型")
	// 加载模型
	ocr.LoadModel()
	fmt.Println("执行识别")
	// 执行识别
	res1 := ocr.Run(rootPath + "/model/demo.png")
	res2 := ocr.Run(rootPath + "/model/demo.jpg")
	fmt.Println(res1 + "\n=====================================\n")
	fmt.Println(res2 + "\n=====================================\n")
}
