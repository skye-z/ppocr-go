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

// 图像识别配置
type OCRConfig struct {
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

// 图像识别模型
type OCRModel struct {
	// 文本检测配置
	DBDetectorOption *C.FD_C_RuntimeOptionWrapper
	// 文本检测模型
	DBDetectorModel *C.FD_C_DBDetectorWrapper
	// 方向分类配置
	ClassifierOption *C.FD_C_RuntimeOptionWrapper
	// 方向分类模型
	ClassifierModel *C.FD_C_ClassifierWrapper
	// 文字识别配置
	RecognizerOption *C.FD_C_RuntimeOptionWrapper
	// 文字识别模型
	RecognizerModel *C.FD_C_RecognizerWrapper
	// 模型整合
	PPOCRModel *C.FD_C_PPOCRv3Wrapper
}

// 识别引擎
type Engine struct {
	Config OCRConfig
	Model  OCRModel
}

// 执行识别
func (e Engine) Run(imgPath string) string {
	var img C.FD_C_Mat = C.FD_C_Imread(C.CString(imgPath))
	var result *C.FD_C_OCRResult = C.FD_C_CreateOCRResult()

	if !e.booleanToGo(C.FD_C_PPOCRv3WrapperPredict(e.Model.PPOCRModel, img, result)) {
		e.destroyOption()
		e.destroyModel()
		C.FD_C_DestroyMat(img)
		C.free(unsafe.Pointer(result))
		return "[Error] Failed to predict"
	}

	var res = (*C.char)(C.malloc(10240))
	defer C.free(unsafe.Pointer(res))

	C.FD_C_OCRResultStr(result, res)

	e.destroyOption()
	e.destroyModel()
	C.FD_C_DestroyOCRResult(result)
	C.FD_C_DestroyMat(img)

	return C.GoString(res)
}

// 加载模型
func (e Engine) LoadModel() {
	// 加载文本检测模型
	var detOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	e.mountOption(detOption)
	var detModel *C.FD_C_DBDetectorWrapper = C.FD_C_CreateDBDetectorWrapper(e.getModel(e.Config.DBDetectorPath), e.getParam(e.Config.DBDetectorPath), detOption, C.FD_C_ModelFormat_PADDLE)
	// 加载方向分类模型
	var clsOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	e.mountOption(clsOption)
	var clsModel *C.FD_C_ClassifierWrapper = C.FD_C_CreateClassifierWrapper(e.getModel(e.Config.ClassifierPath), e.getParam(e.Config.ClassifierPath), clsOption, C.FD_C_ModelFormat_PADDLE)
	// 加载文字识别模型
	var recOption *C.FD_C_RuntimeOptionWrapper = C.FD_C_CreateRuntimeOptionWrapper()
	e.mountOption(recOption)
	var recModel *C.FD_C_RecognizerWrapper = C.FD_C_CreateRecognizerWrapper(e.getModel(e.Config.RecognizerPath), e.getParam(e.Config.RecognizerPath), C.CString(e.Config.KeysPath), recOption, C.FD_C_ModelFormat_PADDLE)
	// 创建PP-OCR
	var ppoceModel *C.FD_C_PPOCRv3Wrapper = C.FD_C_CreatePPOCRv3Wrapper(detModel, clsModel, recModel)

	if !e.booleanToGo(C.FD_C_PPOCRv3WrapperInitialized(ppoceModel)) {
		e.destroyOption()
		e.destroyModel()
		// 初始化失败
		return
	}

	e.Model = OCRModel{
		DBDetectorOption: detOption,
		DBDetectorModel:  detModel,
		ClassifierOption: clsOption,
		ClassifierModel:  clsModel,
		RecognizerOption: recOption,
		RecognizerModel:  recModel,
		PPOCRModel:       ppoceModel,
	}
}

// 挂载模型配置
func (e Engine) mountOption(option *C.FD_C_RuntimeOptionWrapper) {
	if e.Config.UseGPU {
		C.FD_C_RuntimeOptionWrapperUseGpu(option, 0)
	} else {
		C.FD_C_RuntimeOptionWrapperUseCpu(option)
	}
}

// 获取模型地址
func (e Engine) getModel(path string) *C.char {
	return C.CString(path + "/inference.pdmodel")
}

// 获取模型参数地址
func (e Engine) getParam(path string) *C.char {
	return C.CString(path + "/inference.pdiparams")
}

// 销毁模型配置
func (e Engine) destroyOption() {
	C.FD_C_DestroyRuntimeOptionWrapper(e.Model.DBDetectorOption)
	C.FD_C_DestroyRuntimeOptionWrapper(e.Model.ClassifierOption)
	C.FD_C_DestroyRuntimeOptionWrapper(e.Model.RecognizerOption)
}

// 销毁模型本体
func (e Engine) destroyModel() {
	C.FD_C_DestroyDBDetectorWrapper(e.Model.DBDetectorModel)
	C.FD_C_DestroyClassifierWrapper(e.Model.ClassifierModel)
	C.FD_C_DestroyRecognizerWrapper(e.Model.RecognizerModel)
	C.FD_C_DestroyPPOCRv3Wrapper(e.Model.PPOCRModel)
}

// C布尔转Go
func (e Engine) booleanToGo(b C.FD_C_Bool) bool {
	var cFalse C.FD_C_Bool
	return b != cFalse
}
