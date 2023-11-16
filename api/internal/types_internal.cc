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

#include "api/internal/types_internal.h"

namespace fastdeploy {

std::unique_ptr<fastdeploy::RuntimeOption>&
FD_C_CheckAndConvertRuntimeOptionWrapper(
    FD_C_RuntimeOptionWrapper* fd_c_runtime_option_wrapper) {
  FDASSERT(fd_c_runtime_option_wrapper != nullptr,
           "The pointer of fd_c_runtime_option_wrapper shouldn't be nullptr.");
  return fd_c_runtime_option_wrapper->runtime_option;
}

}  // namespace fastdeploy