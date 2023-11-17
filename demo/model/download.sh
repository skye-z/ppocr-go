#!/bin/bash

# 模型与字典下载地址
path_det="https://paddleocr.bj.bcebos.com/PP-OCRv4/chinese/"
name_det="ch_PP-OCRv4_det_infer"
path_cls="https://paddleocr.bj.bcebos.com/dygraph_v2.0/ch/"
name_cls="ch_ppocr_mobile_v2.0_cls_infer"
path_rec="https://paddleocr.bj.bcebos.com/PP-OCRv4/chinese/"
name_rec="ch_PP-OCRv4_rec_infer"
path_key="https://bj.bcebos.com/paddlehub/fastdeploy/"
name_key="ppocr_keys_v1"
# 下载模型与字典
urls=(
    "$path_det$name_det.tar"
    "$path_cls$name_cls.tar"
    "$path_rec$name_rec.tar"
    "$path_key$name_key.txt"
)
echo "Downloading..."
for url in "${urls[@]}"; do
    curl -O "$url" &
done
wait
# 检查下载是否成功
if [ $? -eq 0 ]; then
    echo "Download successful."
    # 解压缩到目标目录
    echo "Extracting Model..."
    tar -zxvf "./$name_det.tar" -C "."
    mv "./$name_det" "./det-v4"
    tar -zxvf "./$name_cls.tar" -C "."
    mv "./$name_cls" "./cls-v2"
    tar -zxvf "./$name_rec.tar" -C "."
    mv "./$name_rec" "./rec-v4"
    mv "./$name_key.txt" "./keys.txt"
    # 检查解压是否成功
    if [ $? -eq 0 ]; then
        rm "./$name_det.tar"
        rm "./$name_cls.tar"
        rm "./$name_rec.tar"
        echo "Extraction successful."
    else
        echo "Extraction failed."
    fi
else
    echo "Download failed."
fi
