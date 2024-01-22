---
layout: post
title: Convert PDF เป็นสีขาวดำบน Ubuntu ด้วย imagemagick
tags: pdf, converter, imagemagick
date: 2024-01-22 07:48 +0700
---
ใช้คำสั่ง

```shell
convert -monochrome -density 600 input.pdf output.pdf
```

ถ้าติดปัญหา `not authorize .....`  หรือ `not allowed by security policy` ให้เข้าไปแก้
security policy ของ PDF ก่อน

```shell
# หา path ของไฟล์ policy ด้วยคำสั่ง -list policy
identify -list policy | less

# Output:
# Path: /etc/ImageMagick-6/policy.xml
#   Policy: Resource
#     name: disk
#     value: 1GiB
#   Policy: Resource
#     name: map
#     value: 512MiB
#   Policy: Resource
#     name: memory
#     value: 256MiB
#   Policy: Resource
#   ...
```

เข้าไป comment หรือลบบรรทัดที่เป็น policy ของ PDF ออก

```xml
<!-- <policy domain="coder" rights="none" pattern="PDF" /> -->
```
