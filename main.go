package main

/*
#cgo CFLAGS: -I./api
#cgo LDFLAGS: -L./lib -lfastdeploy -Wl,-rpath,'${SRCDIR}/lib'
#include <api/vision.h>
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func FDBooleanToGo(b C.FD_C_Bool) bool {
	var cFalse C.FD_C_Bool
	return b != cFalse
}

func UseCPU(detPath, clsPath, recPath, keysPath, imgPath string) string {
	// 加载文本检测模型
	var detOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	C.FD_C_RuntimeOptionWrapperUseCpu(detOption)
	var detModel *C.FD_C_DBDetectorWrapper = C.FD_C_CreateDBDetectorWrapper(C.CString(detPath+"/inference.pdmodel"), C.CString(detPath+"/inference.pdiparams"), detOption, C.FD_C_ModelFormat_PADDLE)
	// 加载方向分类模型
	var clsOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	C.FD_C_RuntimeOptionWrapperUseCpu(clsOption)
	var clsModel *C.FD_C_ClassifierWrapper = C.FD_C_CreateClassifierWrapper(C.CString(clsPath+"/inference.pdmodel"), C.CString(clsPath+"/inference.pdiparams"), clsOption, C.FD_C_ModelFormat_PADDLE)
	// 加载文字识别模型
	var recOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	C.FD_C_RuntimeOptionWrapperUseCpu(recOption)
	var recModel *C.FD_C_RecognizerWrapper = C.FD_C_CreateRecognizerWrapper(C.CString(recPath+"/inference.pdmodel"), C.CString(recPath+"/inference.pdiparams"), C.CString(keysPath), recOption, C.FD_C_ModelFormat_PADDLE)
	// 创建PP-OCR
	var model *C.FD_C_PPOCRv3Wrapper = C.FD_C_CreatePPOCRv3Wrapper(detModel, clsModel, recModel)

	if !FDBooleanToGo(C.FD_C_PPOCRv3WrapperInitialized(model)) {
		C.FD_C_DestroyRuntimeOptionWrapper(detOption)
		C.FD_C_DestroyRuntimeOptionWrapper(clsOption)
		C.FD_C_DestroyRuntimeOptionWrapper(recOption)
		C.FD_C_DestroyDBDetectorWrapper(detModel)
		C.FD_C_DestroyClassifierWrapper(clsModel)
		C.FD_C_DestroyRecognizerWrapper(recModel)
		C.FD_C_DestroyPPOCRv3Wrapper(model)
		return "[Error] Failed to initialize"
	}

	var img C.FD_C_Mat = C.FD_C_Imread(C.CString(imgPath))
	var result *C.FD_C_OCRResult = C.FD_C_CreateOCRResult()

	if !FDBooleanToGo(C.FD_C_PPOCRv3WrapperPredict(model, img, result)) {
		C.FD_C_DestroyRuntimeOptionWrapper(detOption)
		C.FD_C_DestroyRuntimeOptionWrapper(clsOption)
		C.FD_C_DestroyRuntimeOptionWrapper(recOption)
		C.FD_C_DestroyDBDetectorWrapper(detModel)
		C.FD_C_DestroyClassifierWrapper(clsModel)
		C.FD_C_DestroyRecognizerWrapper(recModel)
		C.FD_C_DestroyPPOCRv3Wrapper(model)
		C.FD_C_DestroyMat(img)
		C.free(unsafe.Pointer(result))
		return "[Error] Failed to predict"
	}

	var res = (*C.char)(C.malloc(10240))
	defer C.free(unsafe.Pointer(res))

	C.FD_C_OCRResultStr(result, res)

	C.FD_C_DestroyRuntimeOptionWrapper(detOption)
	C.FD_C_DestroyRuntimeOptionWrapper(clsOption)
	C.FD_C_DestroyRuntimeOptionWrapper(recOption)
	C.FD_C_DestroyDBDetectorWrapper(detModel)
	C.FD_C_DestroyClassifierWrapper(clsModel)
	C.FD_C_DestroyRecognizerWrapper(recModel)
	C.FD_C_DestroyPPOCRv3Wrapper(model)
	C.FD_C_DestroyOCRResult(result)
	C.FD_C_DestroyMat(img)

	return C.GoString(res)
}

func main() {
	rootPath := "/Users/zhaoguiyang/Desktop/Workspace/Golang/ppocr-go/model"
	detPath := rootPath + "/det-v4"
	clsPath := rootPath + "/cls-v2"
	recPath := rootPath + "/rec-v4"
	keysPath := rootPath + "/keys.txt"
	imgPath := rootPath + "/demo.png"
	result := UseCPU(detPath, clsPath, recPath, keysPath, imgPath)

	fmt.Printf("Result: " + result)
}
