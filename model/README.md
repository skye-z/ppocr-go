# Paddle OCR Model

欢迎使用`PPOCR-GO`, 本目录用于存放`Paddle OCR`模型和字典文件, 点此[查看完整模型列表](https://github.com/PaddlePaddle/PaddleOCR/blob/release/2.7/doc/doc_ch/models_list.md).

运行所需文件列表:
1. 检测模型: [v4](https://paddleocr.bj.bcebos.com/PP-OCRv4/chinese/ch_PP-OCRv4_det_infer.tar)、[v3](https://paddleocr.bj.bcebos.com/PP-OCRv3/chinese/ch_PP-OCRv3_det_infer.tar)
2. 方向分类器: [v2](https://paddleocr.bj.bcebos.com/dygraph_v2.0/ch/ch_ppocr_mobile_v2.0_cls_infer.tar)
3. 识别模型: [v4](https://paddleocr.bj.bcebos.com/PP-OCRv4/chinese/ch_PP-OCRv4_rec_infer.tar)、[v3](https://paddleocr.bj.bcebos.com/PP-OCRv3/chinese/ch_PP-OCRv3_rec_infer.tar)
4. 字典: [v1](https://bj.bcebos.com/paddlehub/fastdeploy/ppocr_keys_v1.txt)

> 以上各个文件版本上的差距主要体现在识别精度

你可以根据上述信息自行下载模型后解压到本目录中, 也可使用目录下的`download.sh`脚本自动下载

```shell
# 请cd到本目录后执行
bash download.sh
```