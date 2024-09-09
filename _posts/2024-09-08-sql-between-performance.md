---
layout: post
title: ลองเทียบ performance ของการใช้ฟังก์ชัน COALESCE กับ IS NULL ใน SQL
tags: sql
date: 2024-09-08 16:24 +0700
---

วันก่อนเห็นน้องในทีมส่งข้อความหากันเองในกลุ่มที่ทำงานเป็น query ด้วย `Gorm` หน้าตาประมาณ

```go
query = query.Where("? BETWEEN c.effective_at AND COALESCE(c.expire_at,  ?)", effected_time, effected_time.AddDate(0, 0, 1))
```

พอเห็นแล้วอดไม่ได้ที่จะแนะนำว่าเช็คถ้า `IS NULL` ก่อนก็สามารถทำได้เหมือนกัน ถ้าอยากได้ query ที่รองรับ
เผื่อกรณีที่ `expire_at` สามารถเป็น `NULL` ได้แบบนี้

```go
query = query.Where("c.effective_at <= ? AND (c.expire_at IS NULL OR c.expire_at > ?)", effected_time, effected_time)
```

เพราะในหัวตอนนั้น เข้าใจว่าการใช้ ฟังก์ชันแบบ `COALESCE` น่าจะได้ performance
ที่แย่กว่าเพราะถือเป็นค่าใหม่ ที่ไม่ได้ถูกทำ index เอาไว้

ก็เลยได้ลองเขียน script เก็บไว้ดูเลยว่ามันจะดีกว่าขนาดไหน แต่เท่าที่ดูจากผลลัพธ์ก็ได้ค่าเฉลี่ยที่ไม่ต่างกันเท่าไหร่ ซึ่งอาจจะเพราะ scale ยังเป็นแค่ db ขนาดเล็กด้วย โดยเทียบผลลัพธ์ทั้งแบบที่มี index และไม่มี index บน SQLite3

```sh
# no index
Query 1 - BETWEEN: 68.308µs, IS NULL/OR: 36.579µs
Query 2 - BETWEEN: 25.68µs, IS NULL/OR: 22.227µs
Query 3 - BETWEEN: 26.301µs, IS NULL/OR: 23.961µs
Query 4 - BETWEEN: 25.267µs, IS NULL/OR: 23.52µs

-> BETWEEN Query - Avg: 24.118µs, Max: 68.308µs, Min: 22.626µs
-> IS NULL/OR Query - Avg: 21.933µs, Max: 28.717µs, Min: 20.758µs
```

```sh
#  index expire_at
Query 1 - BETWEEN: 35.829µs, IS NULL/OR: 34.539µs
Query 2 - BETWEEN: 20.68µs, IS NULL/OR: 20.437µs
Query 3 - BETWEEN: 16.351µs, IS NULL/OR: 17.691µs
Query 4 - BETWEEN: 15.277µs, IS NULL/OR: 16.92µs
.
.
Query 57 - BETWEEN: 14.785µs, IS NULL/OR: 16.027µs
Query 58 - BETWEEN: 14.753µs, IS NULL/OR: 17.613µs
Query 59 - BETWEEN: 14.972µs, IS NULL/OR: 16.867µs
Query 60 - BETWEEN: 15.153µs, IS NULL/OR: 16.241µs
Query 61 - BETWEEN: 14.869µs, IS NULL/OR: 16.411µs
Query 62 - BETWEEN: 14.944µs, IS NULL/OR: 16.131µs
Query 63 - BETWEEN: 14.735µs, IS NULL/OR: 16.196µs
Query 64 - BETWEEN: 15.084µs, IS NULL/OR: 17.143µs
Query 65 - BETWEEN: 15.457µs, IS NULL/OR: 16.688µs
.
.
Query 80 - BETWEEN: 25.236µs, IS NULL/OR: 22.812µs
Query 81 - BETWEEN: 24.781µs, IS NULL/OR: 22.533µs
Query 82 - BETWEEN: 24.459µs, IS NULL/OR: 22.659µs
Query 83 - BETWEEN: 24.48µs, IS NULL/OR: 24.181µs
Query 84 - BETWEEN: 24.922µs, IS NULL/OR: 23.309µs
Query 85 - BETWEEN: 24.829µs, IS NULL/OR: 22.028µs
Query 86 - BETWEEN: 24.633µs, IS NULL/OR: 22.383µs

-> BETWEEN Query - Avg: 20.556µs, Max: 130.997µs, Min: 14.37µs
-> IS NULL/OR Query - Avg: 19.853µs, Max: 34.539µs, Min: 13.835µs

```

```sh
#  index effective_at, expire_at
Query 1 - BETWEEN: 64.32µs, IS NULL/OR: 26.332µs
Query 2 - BETWEEN: 31.684µs, IS NULL/OR: 21.047µs
Query 3 - BETWEEN: 27.349µs, IS NULL/OR: 18.24µs
.
Query 90 - BETWEEN: 14.501µs, IS NULL/OR: 13.311µs
Query 91 - BETWEEN: 14.412µs, IS NULL/OR: 14.461µs
Query 92 - BETWEEN: 14.897µs, IS NULL/OR: 14.527µs
Query 93 - BETWEEN: 14.834µs, IS NULL/OR: 13.487µs
Query 94 - BETWEEN: 14.44µs, IS NULL/OR: 13.501µs
Query 95 - BETWEEN: 14.643µs, IS NULL/OR: 13.488µs
Query 96 - BETWEEN: 14.544µs, IS NULL/OR: 13.745µs
Query 97 - BETWEEN: 14.559µs, IS NULL/OR: 13.499µs
Query 98 - BETWEEN: 14.548µs, IS NULL/OR: 13.602µs
Query 99 - BETWEEN: 14.576µs, IS NULL/OR: 13.42µs
Query 100 - BETWEEN: 14.421µs, IS NULL/OR: 14.638µs

-> BETWEEN Query - Avg: 22.53µs, Max: 64.32µs, Min: 14.355µs
-> IS NULL/OR Query - Avg: 17.132µs, Max: 26.332µs, Min: 13.311µs
```

> script ถูกแชร์ไว้ที่ [github](https://github.com/wittawasw/wittawasw.github.io/blob/main/code/go/sqlperf/main.go)
