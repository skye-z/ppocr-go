// Copyright (c) 2023 PaddlePaddle Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#pragma once

#include "api/core/fd_common.h"
#include "api/core/fd_type.h"
#include "api/runtime/runtime_option.h"
#include "api/vision/result.h"

typedef struct FD_C_PaddleClasModelWrapper FD_C_PaddleClasModelWrapper;

#ifdef __cplusplus
extern "C" {
#endif

/** \brief Create a new FD_C_PaddleClasModelWrapper object
 *
 * \param[in] model_file Path of model file, e.g resnet/model.pdmodel
 * \param[in] params_file Path of parameter file, e.g resnet/model.pdiparams, if the model format is ONNX, this parameter will be ignored
 * \param[in] config_file Path of configuration file for deployment, e.g resnet/infer_cfg.yml
 * \param[in] fd_c_runtime_option_wrapper RuntimeOption for inference, the default will use cpu, and choose the backend defined in `valid_cpu_backends`
 * \param[in] model_format Model format of the loaded model, default is Paddle format
 *
 * \return Return a pointer to FD_C_PaddleClasModelWrapper object
 */

FASTDEPLOY_CAPI_EXPORT extern __fd_give FD_C_PaddleClasModelWrapper*
FD_C_CreatePaddleClasModelWrapper(
    const char* model_file, const char* params_file, const char* config_file,
    FD_C_RuntimeOptionWrapper* fd_c_runtime_option_wrapper,
    const FD_C_ModelFormat model_format);

/** \brief Destroy a FD_C_PaddleClasModelWrapper object
 *
 * \param[in] fd_c_paddleclas_model_wrapper pointer to FD_C_PaddleClasModelWrapper object
 */

FASTDEPLOY_CAPI_EXPORT extern void FD_C_DestroyPaddleClasModelWrapper(
    __fd_take FD_C_PaddleClasModelWrapper* fd_c_paddleclas_model_wrapper);

/** \brief Predict the classification result for an input image
 *
 * \param[in] fd_c_paddleclas_model_wrapper pointer to FD_C_PaddleClasModelWrapper object
 * \param[in] img pointer to cv::Mat image
 * \param[in] fd_c_classify_result pointer to FD_C_ClassifyResult object, which stores the result.
 */

FASTDEPLOY_CAPI_EXPORT extern FD_C_Bool FD_C_PaddleClasModelWrapperPredict(
    __fd_keep FD_C_PaddleClasModelWrapper* fd_c_paddleclas_model_wrapper,
    FD_C_Mat img, FD_C_ClassifyResult* fd_c_classify_result_wrapper);

/** \brief Check if the model is initialized successfully
 *
 * \param[in] fd_c_paddleclas_model_wrapper pointer to FD_C_PaddleClasModelWrapper object
 *
 * \return Return a bool of value true if initialized successfully
 */

FASTDEPLOY_CAPI_EXPORT extern FD_C_Bool FD_C_PaddleClasModelWrapperInitialized(
    __fd_keep FD_C_PaddleClasModelWrapper* fd_c_paddleclas_model_wrapper);


/** \brief Predict the classification results for a batch of input images
   *
   * \param[in] fd_c_paddleclas_model_wrapper pointer to FD_C_PaddleClasModelWrapper object
   * \param[in] imgs The input image list, each element comes from cv::imread()
   * \param[in] results The output classification result list
   * \return true if the prediction successed, otherwise false
   */
FASTDEPLOY_CAPI_EXPORT extern FD_C_Bool FD_C_PaddleClasModelWrapperBatchPredict(
            __fd_keep FD_C_PaddleClasModelWrapper* fd_c_paddleclas_model_wrapper,
            FD_C_OneDimMat imgs,
            FD_C_OneDimClassifyResult* results);


#ifdef __cplusplus
}  // extern "C"
#endif
