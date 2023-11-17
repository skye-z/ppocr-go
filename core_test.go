package ppocr

import (
	"fmt"
	"os"
	"testing"
)

func TestOCR(t *testing.T) {
	rootPath, _ := os.Getwd()
	config := Config{
		UseGPU:         false,
		DBDetectorPath: rootPath + "/model/det-v4",
		ClassifierPath: rootPath + "/model/cls-v2",
		RecognizerPath: rootPath + "/model/rec-v4",
		KeysPath:       rootPath + "/model/keys.txt",
	}
	res1 := Run(config, rootPath+"/model/demo.png")
	res2 := Run(config, rootPath+"/model/demo.jpg")
	fmt.Println(res1 + "\n=====================================\n")
	fmt.Println(res2 + "\n=====================================\n")
}
