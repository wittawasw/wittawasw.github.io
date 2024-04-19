---
layout: post
title: Rate Limit ใน Ruby on Rails
tags: rails, rate limit
date: 2024-04-19 07:28 +0700
---
Rate Limit เป็นหนึ่งในปัญหาที่ทุก tech stack มีวิธีการแก้ปัญหาที่คล้ายกัน ถึงแม้ว่ารายละเอียดการ
implement จะแตกต่างกันบ้าง แต่ลักษณะการทำงานเบื้องต้นจะคล้ายกันเสมอ คือ

- request จาก client ไปถึง server
- server ทำการนับจำนวน request ที่เค้ามาแต่ล่ะครั้ง โดยเก็บจำนวนของ request ใน datastore
- ตรวจสอบว่า request ที่เข้ามาถ้านับรวมกับจำนวนเดิมที่เก็บ จะเกิน limit ตามที่กำหนดไว้หรือยัง
  โดยอาจกำหนดเป็น limit ของจำนวนของ request ภายในช่วงเวลา เช่น
  100 requests ต่อ 10 นาที
- ปฏิเสธ request ถ้า เกิน limit และยอมให้ผ่านไปถึง resource ได้ถ้าไม่เกิน

### Resource
อาจหมายถึง database โดยตรงหรือ service อื่นที่เราต้องการลดความเสี่ยงที่จะ
ทำงานเกินกำลังและทำให้ระบบเกิด downtime หรือ error อื่นๆ

### Datastore
ส่วนที่ใช้เก็บค่าจำนวน request ซึ่งเป็นได้ทั้งการเก็บโดยตรงภายในหน่วยความจำหรือการเก็บภายใน
ซอฟต์แวร์ฐานข้อมูล ขึ้นอยู่กับลักษณะการใช้งาน ที่หากเป็นการใช้ภายใน process เดียวกัน ก็อาจเก็บค่า
ภายในหน่วยความจำโดยตรง แต่หากมีจำนวน server ที่รองรับหลายอัน แล้วต้องการจำกัดจำนวน limit
ร่วมกันก็ควรใช้ซอฟต์แวร์ฐานข้อมูลช่วย

```
+--------+            +---------+            +---------+
|        |----------->|         |----------->|         |
| Client |            | Server  |            |Resources|
|        |<-----------|         |<-----------|         |
+--------+            +---------+            +---------+
                          |
                          |
                          v
                    +-------------+
                    |  Datastore  |
                    +-------------+
```

<p style="text-align: center;"><em>แบบง่ายที่เป็นการใช้ภายใน process เดียว</em></p>
<p style="text-align: center;"><em>---------------------</em></p>


```
+--------+            +--------+
|        |            |        |
|        |     ------>| Server |<-
|        |    /       |        |  \
|        |   /        +--------+   \
|        |  /         +--------+    \            +---------+
|        |<-          |        |     \           |         |
| Client |<---------->| Server |<--------------->|Resources|
|        |<-          |        |<-     \         |         |
|        |  \         +--------+  \     \        +---------+
|        |   \        +--------+   \     \       +-------------+
|        |    \       |        |<--------------->|  Datastore  |
|        |     ------>| Server |                 +-------------+
|        |            |        |
+--------+            +--------+
```
<p style="text-align: center;"><em>แบบที่ใช้ datastore ร่วมกัน</em></p>

## ความซับซ้อนที่เกิดขึ้นเสมอ

หาก resource ที่เราต้องกันการเข้าถึง มีความต้องการในการจำกัดที่แตกต่างกัน รูปแบบการทำ
rate limit ก็จะเพิ่มความซับซ้อนและความยุ่งยากในการดูแลมากขึ้น แนวคิดและรูปแบบการทำ
rate limit ของ Rails จึงช่วยได้มาก หากไม่มีความเข้าใจ หรือไม่ต้องการ implement ด้วยตัวเอง

## ตัวอย่างการ Rate Limit แยกตาม controller ของ Rails

```ruby

class SessionsController < ApplicationController
  # จำกัด 10 requests ต่อ 3 นาที
  # ที่ method: :create เท่านั้น
  rate_limit to: 10, within: 3.minutes, only: :create
end

class SignupsController < ApplicationController
  # จำกัด 1000 requests ต่อ 10 วินาที
  # แยกตาม domain name ที่ request
  # โดยให้ redirect ไปที่ busy_controller_url พร้อมกับ alert
  # ที่ method: :new เท่านั้น
  rate_limit to: 1000, within: 10.seconds,
    by: -> { request.domain },
    with: -> { redirect_to busy_controller_url, alert: "Too many signups on domain!" },
    only: :new
end

class APIController < ApplicationController
  # จำกัด 10 requests ต่อ ต่อ 3 นาที
  # โดยใช้ datastore เป็น redis ที่ REDIS_URL
  RATE_LIMIT_STORE = ActiveSupport::Cache::RedisCacheStore.new(url: ENV["REDIS_URL"])
  rate_limit to: 10, within: 3.minutes, store: RATE_LIMIT_STORE
end
```

ที่มา: [github.com/rails/rails](https://github.com/rails/rails/blob/main/actionpack/lib/action_controller/metal/rate_limiting.rb)

> ปัจจุบัน ณ วันที่เขียน blog นี้
> feature ยังไม่ได้ถูก release เป็น version แต่ถูกรวมอยู่ใน `main` แล้ว
