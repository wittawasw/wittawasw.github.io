---
layout: post
title: ติดตั้ง GVM กับ Go 1.22 บน MacOS ตั้งแต่ต้น
tags: gvm, golang, install, macos
date: 2024-02-14 10:51 +0700
---
> บทความนี้จะทำการติดตั้ง Go เวอร์ชั่น 1.22 โดยผ่าน [GVM](https://github.com/moovweb/gvm) เพื่อเผื่อทางเลือกสำหรับการติดตั้งเวอร์ชั่นอื่นด้วย

อ้างอิง: 
- [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm)
- [homebrew](https://formulae.brew.sh/)
- [Go](https://go.dev/)

## ติดตั้ง gvm 
```shell
# ความต้องการของ script ติดตั้ง
xcode-select --install
brew update
brew install mercurial

# ติดตั้ง gvm
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

## ติดตั้ง Go
```shell
# เนื่องจาก Go 1.5+ ต้องใช้ Go ในการ compile จึงเลือกใช้การติดตั้งผ่าน Homebrew ก่อน เพื่อความสะดวก
brew install go

# ติดตั้ง Go 1.22
gvm install go1.22

# ลบ Go ที่ติดตั้งผ่าน homebrew ออก
brew uninstall go

# ตั้งค่าให้ใช้ Go 1.22 เป็นค่าเริ่มต้น
gvm use go1.22 --default
```

## ตรวจสอบ Go ที่ติดตั้ง
```shell
go version
# => go version go1.22.0 darwin/arm64
```

