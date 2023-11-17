package main

import (
	"fmt"
	"os"

	ppocr "github.com/skye-z/ppocr-go"
)

func main() {
	rootPath, _ := os.Getwd()
	// 创建引擎
	ocr := ppocr.Engine{
		Config: ppocr.OCRConfig{
			UseGPU:         false,
			DBDetectorPath: rootPath + "/det-v4",
			ClassifierPath: rootPath + "/cls-v2",
			RecognizerPath: rootPath + "/rec-v4",
			KeysPath:       rootPath + "/keys.txt",
		},
	}
	// 加载模型
	ocr.LoadModel()
	// 执行识别
	res1 := ocr.Run(rootPath + "/demo.png")
	res2 := ocr.Run(rootPath + "/demo.jpg")
	fmt.Println(res1 + "\n=====================================\n")
	fmt.Println(res2 + "\n=====================================\n")
}
