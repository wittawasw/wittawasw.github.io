---
layout: post
title: Ruby 3.3.5 เตือนให้เอา standard lib ออก
tags: ruby deprecation
keywords: ruby, deprecation
description: ไม่รู้ว่าเตือนตั้งแต่ 3.3.อะไร เพราะจาก 3.3.0 ก็มา 3.3.5 เลย เจอตอนกำลังจะเขียน
  blog ก็เลยได้อีก blog ว่าตอนนี้ Ruby core กำลังเดินตาม Roadmap ที่อยากให้ Runtime
  โดย default มี memory footprint ที่เล็กลง
date: 2024-09-04 21:34 +0700
---
ไม่รู้ว่าเตือนตั้งแต่ 3.3.อะไร เพราะจาก 3.3.0 ก็มา 3.3.5 เลย เจอตอนกำลังจะเขียน
blog ก็เลยได้อีก blog ว่าตอนนี้ Ruby core กำลังเดินตาม Roadmap ที่อยากให้
Runtime โดย default มี memory footprint ที่เล็กลง โดยเท่าที่รู้ก็ถอดออกไปตามนี้

```ruby
# Ruby 3.5 deprecation warnings
gem "logger"
gem "base64"
gem "csv"
gem "ostruct"
```

ก็ไม่ใช่เรื่องแย่ แต่ก็น่าเสียดายที่คำสั่งใน `base64` เป็น 1 ใน คำสั่งแสนสะดวก
ที่ผมเอาไว้ใช้ได้ตลอดต้องมาโดนถอดออก แต่ `ostruct` ไม่น่าเสียดายเลยเพราะเอาจริงๆ
มันถูกมองว่าไม่ควรใช้มานานแล้ว
