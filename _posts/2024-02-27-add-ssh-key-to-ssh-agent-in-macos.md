---
layout: post
title: เพิ่ม SSH key ให้ ssh-agent ใน MacOS
tags: macos, ssh
date: 2024-02-27 09:59 +0700
---
นอกจากขั้นตอนปกติที่เราแก้ใน file `~/.ssh/config` แล้ว การเพิ่ม SSH key ที่สร้างใหม่ให้สามารถใช้งานได้ ใน MacOS ต้องทำการเพิ่มเข้า Apple Keychain ด้วยคำสั่ง

```shell
ssh-add --apple-use-keychain ~/.ssh/wittawasw_rsa
```
