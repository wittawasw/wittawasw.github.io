---
layout: post
title: Install torch-rb บน Ubuntu
tags: ruby ml
keywords: ruby, ml
description: ''
date: 2024-10-08 12:18 +0700
---
แก้ไขเพิ่มเติมจาก [readme ของ torch-rb](https://github.com/ankane/torch.rb)

## ดาวน์โหลดไฟล์ libtorch

```sh
# หา version ล่าสุดจากในนี้
# https://download.pytorch.org/libtorch/cpu/
# ใช้ cxxx-11 สำหรับ CPU
# กับเวอร์ชั่นล่าสุดตอนที่เขียนคือ 2.4.1
curl -L https://download.pytorch.org/libtorch/cpu/libtorch-cxx11-abi-shared-with-deps-2.4.1%2Bcpu.zip > libtorch.zip

unzip -q libtorch.zip
# rm libtorch.zip # Optional
```

## Install ผ่าน Gemfile

```sh
bundle config build.torch-rb --with-torch-dir=/path/to/libtorch
bundle add torch-rb

# สามารถใช้ $(pwd)/relative_path/to/libtorch กับ --with-torch-dir ได้
# เช่น bundle config build.torch-rb --with-torch-dir=$(pwd)/libtorch
# หรือตั้งค่าเป็น global ด้วย
# bundle config --global build.torch-rb --with-torch-dir=$(pwd)/libtorch
```
