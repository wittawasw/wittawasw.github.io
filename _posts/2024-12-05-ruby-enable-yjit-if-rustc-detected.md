---
layout: post
title: ค่าตั้งต้นของ Ruby จะใส่ YJIT มาให้ถ้าในเครื่องมี Rust
tags: Ruby Rust YJIT
keywords: Ruby, Rust, YJIT
description:
date: 2024-12-05 18:22 +0700
---

เร็วๆนี้มีกระแสความสนใจในภาษา Rust ในกลุ่มคนไทยมากขึ้น
หลังจากที่มีโอกาสได้แลกเปลี่ยนกับหลายคนก็รู้สึกว่าเราไม่ค่อยรู้ตัวกันเท่าไหร่
ว่าทุกวันนี้ Rust มันแทรกซึมเข้าไปใน tech stack ทุกที่แล้ว อย่างเช่นใน
Ruby ปัจจุบันก็จะ compile `YJIT` ให้ทันทีถ้าตอน install ตรวจพบว่าในเครื่องมี
`rustc` ลงอยู่แล้ว

## วิธีตรวจว่า Ruby รองรับ YJIT

```sh
# สังเกต YJIT ที่ต่อท้ายเลขเวอร์ชัน
ruby -v
# => ruby 3.3.5 (2024-09-03 revision ef084cc8f4) +YJIT [arm64-darwin23]

# แต่ต่อให้ลงให้แล้วก็จะยังไม่เปิดใช้ทันทีเป็นค่าตั้งต้น
ruby -e "puts RubyVM::YJIT.enabled?"
# => false

# ถ้าอยากใช้ต้อง flag บอกด้วย --yjit
ruby --yjit -e "puts RubyVM::YJIT.enabled?"
# => true

# หรือตั้ง env ก่อน
export RUBY_YJIT_ENABLE=1
ruby -e "puts RubyVM::YJIT.enabled?"
# => true
```
