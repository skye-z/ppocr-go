package ppocr

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
	"unsafe"
)

// 识别引擎
type Config struct {
	// 使用显卡
	UseGPU bool
	// 文本检测模型地址
	DBDetectorPath string
	// 方向分类模型地址
	ClassifierPath string
	// 文字识别模型地址
	RecognizerPath string
	// 文字识别字典地址
	KeysPath string
}

// 执行识别
func Run(config Config, imgPath string) string {
	// 加载文本检测模型
	var detOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	mountOption(config.UseGPU, detOption)
	var detModel *C.FD_C_DBDetectorWrapper = C.FD_C_CreateDBDetectorWrapper(getModel(config.DBDetectorPath), getParam(config.DBDetectorPath), detOption, C.FD_C_ModelFormat_PADDLE)
	// 加载方向分类模型
	var clsOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	mountOption(config.UseGPU, clsOption)
	var clsModel *C.FD_C_ClassifierWrapper = C.FD_C_CreateClassifierWrapper(getModel(config.ClassifierPath), getParam(config.ClassifierPath), clsOption, C.FD_C_ModelFormat_PADDLE)
	// 加载文字识别模型
	var recOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	mountOption(config.UseGPU, recOption)
	var recModel *C.FD_C_RecognizerWrapper = C.FD_C_CreateRecognizerWrapper(getModel(config.RecognizerPath), getParam(config.RecognizerPath), C.CString(config.KeysPath), recOption, C.FD_C_ModelFormat_PADDLE)
	// 创建PP-OCR
	var ppoceModel *C.FD_C_PPOCRv3Wrapper = C.FD_C_CreatePPOCRv3Wrapper(detModel, clsModel, recModel)

	if !booleanToGo(C.FD_C_PPOCRv3WrapperInitialized(ppoceModel)) {
		destroyOption(detOption, clsOption, recOption)
		destroyModel(detModel, clsModel, recModel, ppoceModel)
		// 初始化失败
		return "初始化失败"
	}

	var img C.FD_C_Mat = C.FD_C_Imread(C.CString(imgPath))
	var result *C.FD_C_OCRResult = C.FD_C_CreateOCRResult()

	if !booleanToGo(C.FD_C_PPOCRv3WrapperPredict(ppoceModel, img, result)) {
		destroyOption(detOption, clsOption, recOption)
		destroyModel(detModel, clsModel, recModel, ppoceModel)
		C.FD_C_DestroyMat(img)
		C.free(unsafe.Pointer(result))
		return "[Error] Failed to predict"
	}
	var res = (*C.char)(C.malloc(10240))
	defer C.free(unsafe.Pointer(res))

	C.FD_C_OCRResultStr(result, res)

	destroyOption(detOption, clsOption, recOption)
	destroyModel(detModel, clsModel, recModel, ppoceModel)
	C.FD_C_DestroyOCRResult(result)
	C.FD_C_DestroyMat(img)
	// 返回识别结果
	return C.GoString(res)
}

// 挂载模型配置
func mountOption(useGPU bool, option *C.FD_C_RuntimeOptionWrapper) {
	if useGPU {
		C.FD_C_RuntimeOptionWrapperUseGpu(option, 0)
	} else {
		C.FD_C_RuntimeOptionWrapperUseCpu(option)
	}
}

// 获取模型地址
func getModel(path string) *C.char {
	return C.CString(path + "/inference.pdmodel")
}

// 获取模型参数地址
func getParam(path string) *C.char {
	return C.CString(path + "/inference.pdiparams")
}

// 销毁模型配置
func destroyOption(det *C.FD_C_RuntimeOptionWrapper, cls *C.FD_C_RuntimeOptionWrapper, rec *C.FD_C_RuntimeOptionWrapper) {
	C.FD_C_DestroyRuntimeOptionWrapper(det)
	C.FD_C_DestroyRuntimeOptionWrapper(cls)
	C.FD_C_DestroyRuntimeOptionWrapper(rec)
}

// 销毁模型本体
func destroyModel(det *C.FD_C_DBDetectorWrapper, cls *C.FD_C_ClassifierWrapper, rec *C.FD_C_RecognizerWrapper, model *C.FD_C_PPOCRv3Wrapper) {
	C.FD_C_DestroyDBDetectorWrapper(det)
	C.FD_C_DestroyClassifierWrapper(cls)
	C.FD_C_DestroyRecognizerWrapper(rec)
	C.FD_C_DestroyPPOCRv3Wrapper(model)
}

// C布尔转Go
func booleanToGo(b C.FD_C_Bool) bool {
	var cFalse C.FD_C_Bool
	return b != cFalse
}
