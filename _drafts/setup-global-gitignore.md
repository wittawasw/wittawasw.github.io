---
layout: post
title: การตั้งค่า .gitignore สำหรับทุก git directory ในเครื่อง
tags: devtool git
keywords: devtool, git
description: ''
---

บันทึกไว้ว่าในเครื่องมี `gitignore` อะไรบ้าง

## คำสั่งที่ใช้

```sh
# สั่งให้ใช้ไฟล์ ~/.gitignore เป็น reference
git config --global core.excludesFile '~/.gitignore'
```

## ไฟล์ `.gitignore`

```sh
# Environment Variables
.env
.rbenv-vars

# Python
__pycache__

# Node
node_modules
npm-debug.log

# Dart, FLutter
**/doc/api/
**/ios/Flutter/.last_build_id
.dart_tool/
.flutter-plugins
.flutter-plugins-dependencies
.pub-cache/
.pub/
/build/

# Editor
*.iml
*.ipr
*.iws
.idea/

# scratchpad เอาไว้ note สิ่งที่ต้องทำ, สิ่งที่กำลังทำ , ฯลฯ
todo
todo.txt

# Mac
.DS_Store

# Windows
Thumbs.db

```
