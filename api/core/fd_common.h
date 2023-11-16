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

#include <stdint.h>
#include <stdio.h>

#if defined(_WIN32)
#ifdef FD_CAPI
#define FASTDEPLOY_CAPI_EXPORT __declspec(dllexport)
#else
#define FASTDEPLOY_CAPI_EXPORT __declspec(dllimport)
#endif  // FD_CAPI
#else
#define FASTDEPLOY_CAPI_EXPORT __attribute__((visibility("default")))
#endif  // _WIN32

///
/// __fd_give means that a new object is returned. The user should make sure
/// that the returned pointer is used exactly once as a value for an __fd_take
/// argument. In between, it can be used as a value for as many __fd_keep
/// arguments as the user likes.
///
#ifndef __fd_give
#define __fd_give
#endif
///
/// __fd_take means that the object the argument points to is taken over by the
/// function and may no longer be used by the user as an argument to any other
/// function. The pointer value must be one returned by a function returning an
/// __fd_give pointer.
///
#ifndef __fd_take
#define __fd_take
#endif
///
/// __fd_keep means that the function will only use the object temporarily. The
/// object which the argument points to is not taken over by the function. After
/// the function has finished, the user can still use it as an argument to other
/// functions.
///
#ifndef __fd_keep
#define __fd_keep
#endif

typedef int8_t FD_C_Bool;
#define TRUE 1
#define FALSE 0
