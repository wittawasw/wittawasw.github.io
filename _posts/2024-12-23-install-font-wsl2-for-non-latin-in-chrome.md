---
layout: post
title: Install font ใน wsl2 เพื่อให้ Chrome สามารถแสดง Non-latin ได้
tags: wsl2 chrome font
keywords: wsl2, chrome, font
description: ไม่เจอปัญหานี้ใน MacOS แต่เจอใน WSL2 ที่เป็น Ubuntu 24.04 เลยลอง install
  package Noto ดูเพื่อทดสอบว่าจะทำให้แสดงได้หรือไม่ ซึ่งก็ทำงานได้ดี
date: 2024-12-23 11:55 +0700
---
ไม่เจอปัญหานี้ใน MacOS แต่เจอใน WSL2 ที่เป็น Ubuntu 24.04 เลยลอง install
package Noto ดูเพื่อทดสอบว่าจะทำให้แสดงได้หรือไม่ ซึ่งก็ทำงานได้ดี

## วิธี install

```sh
sudo apt install -y fonts-noto-core fonts-noto-cjk fonts-noto-color-emoji

# fonts-noto-core - ฟอนต์หลัก
# fonts-noto-cjk - สำหรับจีน เกาหลี ญี่ปุ่น
# fonts-noto-color-emoji - สำหรับ emoji
```

## วิธีทดสอบ

```sh
# กรณีใช้ google-chrome command
google-chrome --headless --disable-gpu --screenshot --virtual-time-budget=10000 --no-sandbox https://wittawasw.com

# ใช้ ronin-web
ronin-web screenshot --browser google-chrome https://www.wittawasw.com
```
