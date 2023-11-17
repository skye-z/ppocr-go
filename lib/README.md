# FastDeploy Library

欢迎使用`PPOCR-GO`, 本目录用于存放`FastDeploy`的动态链接库.

你可以使用下方列出的预编译库

FastDeploy SDK 1.0.7:
* GPU: 
    * Linux: [x64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-x64-gpu-1.0.7.tgz) 
    * Windows: [x64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-win-x64-gpu-1.0.7.zip) 
* CPU: 
    * Linux: [x64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-x64-1.0.7.tgz)、[aarch64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-aarch64-1.0.7.tgz)
    * Windows: [x64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-win-x64-1.0.7.zip)
    * MacOS: [x64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-osx-x86_64-1.0.7.tgz)、[arm64](https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-osx-arm64-1.0.7.tgz)

如果上述平台没有你当前使用的平台, 你也可点此查看[FastDeploy安装说明](https://github.com/PaddlePaddle/FastDeploy/tree/develop/docs/cn/build_and_install)进行手动编译.

在下载完成后, 请解压后将`third_libs/install`下的所有`.dylib/.so/.lib`文件以及`lib`目录下的所有文件移动到本目录下.

如果你不想这么麻烦, 也可使用目录下的`download.sh`脚本自动下载并移动库文件

```shell
# 请cd到本目录后执行
bash download.sh
```

> 重要提示: 不要在生产环境下载动态链接库, 请在开发环境下载调试好后和整个项目一起部署过去

### 已知问题

1. opencv存在版本号错误问题: 查看报错信息中的版本号, 将本目录下带opencv的文件全部改为这个版本号