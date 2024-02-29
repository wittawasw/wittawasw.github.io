---
layout: post
title: เพิ่ม SSH key ให้ ssh-agent ใน MacOS
tags: macos, ssh
date: 2024-02-27 09:59 +0700
---

เพิ่งค้นพบว่า MacOS หลังกลับมาใช้ไม่นาน ว่าไม่เรียกใช้ SSH key ในชื่ออื่นที่ไม่ใช้การใช้ `id_` นำหน้าแล้วตามด้วย algorithm ที่ใช้เข้ารหัส เช่น `id_rsa`, `id_ed25519` ตั้งแต่ต้น

แต่เนื่องจากติดนิสัยในการสร้าง key ตามงานที่ใช้มาตลอดตั้งแต่ที่ใช้ Ubuntu ว่าจะไม่ใช้ key ซ้ำกันในแต่ล่ะงาน ก็เลยมีปัญหาเล็กน้อยก่อนเริ่มงานล่าสุด เช่น ถ้าชื่องาน `wittawasw` ก็จะตั้งชื่อ key เป็น `wittawasw_rsa` หรือ `wittawasw_ed25519`

### ทางแก้: ใช้คำสั่งเพิ่ม SSH key ให้ ssh-agent โดยตรง

นอกจากขั้นตอนปกติที่เราแก้ใน file `~/.ssh/config` แล้ว การเพิ่ม SSH key ที่สร้างใหม่ให้สามารถใช้งานได้ ใน MacOS ต้องทำการเพิ่มเข้า Apple Keychain ด้วยคำสั่ง

```shell
ssh-add --apple-use-keychain ~/.ssh/wittawasw_rsa
```
