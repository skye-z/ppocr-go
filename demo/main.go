package main

import (
	"fmt"
	"os"

	ppocr "github.com/skye-z/ppocr-go"
)

func main() {
	rootPath, _ := os.Getwd()
	config := ppocr.Config{
		UseGPU:         false,
		DBDetectorPath: rootPath + "/model/det-v4",
		ClassifierPath: rootPath + "/model/cls-v2",
		RecognizerPath: rootPath + "/model/rec-v4",
		KeysPath:       rootPath + "/model/keys.txt",
	}
	res1 := ppocr.Run(config, rootPath+"/model/demo.png")
	res2 := ppocr.Run(config, rootPath+"/model/demo.jpg")
	fmt.Println(res1 + "\n=====================================\n")
	fmt.Println(res2 + "\n=====================================\n")
}
