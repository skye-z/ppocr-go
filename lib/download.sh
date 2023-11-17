#!/bin/bash

echo "Please select a platform:"
echo "1. GPU Linux x64"
echo "2. GPU Windows x64"
echo "3. CPU Linux x64"
echo "4. CPU Linux aarch64"
echo "5. CPU Windows x64"
echo "6. CPU MacOS x64"
echo "7. CPU MacOS arm64"
read -p "Please enter the option number: " choice

case $choice in
    1)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-x64-gpu-1.0.7.tgz"
        type=".so"
        ;;
    2)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-win-x64-gpu-1.0.7.zip"
        type=".lib"
        ;;
    3)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-x64-1.0.7.tgz"
        type=".so"
        ;;
    4)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-linux-aarch64-1.0.7.tgz"
        type=".so"
        ;;
    5)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-win-x64-1.0.7.zip"
        type=".lib"
        ;;
    6)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-osx-x86_64-1.0.7.tgz"
        type=".dylib"
        ;;
    7)
        url="https://bj.bcebos.com/fastdeploy/release/cpp/fastdeploy-osx-arm64-1.0.7.tgz"
        type=".dylib"
        ;;
    *)
        echo "无效的选项"
        exit 1
        ;;
esac

echo "Downloading..."
curl -O $url
filename=$(basename $url)
echo "Download successful."
mkdir -p $PWD/cache
tar -zxvf $filename -C $PWD/cache
echo "Extraction successful."
rm $filename

# 移动.dylib、.so和.lib文件
find ./cache -maxdepth 3 -path "./cache/fastdeploy-*/lib/*" -exec mv {} . \;
find ./cache -type f \( -path "./cache/fastdeploy-*/third_libs/install/*$type" -o -path "./cache/fastdeploy-*/third_libs/install/openvino/runtime/lib/plugins.xml" \) -exec mv {} . \;
wait
rm -rf ./cache
