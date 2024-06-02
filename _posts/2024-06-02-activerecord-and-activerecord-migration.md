---
layout: post
title: เรื่องที่รู้ไว้ก็ดีก่อนเริ่มใช้งาน Active Record และ Active Record Migration
tags: rails database
keywords: rails, database
description: 'ลองพยายามเขียนอะไรซักอย่างขึ้นมาอุดช่องว่าง
สำหรับคนที่พอเข้าใจการทำงานของฐานข้อมูลอยู่แล้ว แต่แค่ไม่ได้ทำงานด้วย Ruby on Rails มาก่อน
ให้สามารถเห็นภาพการทำงานที่ชัดขึ้น'
date: 2024-06-02 14:28 +0700
---

บทความนี้ได้แรงผลักดันมาจากการที่โดนถามมาเรื่องการทำ Database Migration ของ Rails
ซึ่งส่วนตัวคิดว่าบทความในอินเตอร์เนตเท่าที่หาดู ไม่มีอันไหนช่วยให้เข้าใจได้ดีกว่า
[Official Doc](https://guides.rubyonrails.org/active_record_migrations.html) ของ Rails เลย
แต่เราก็เชื่อว่าคนถามน่าจะลองอ่านแล้วล่ะ แต่แค่อาจจะยังไม่ตกผลึกออกมาเป็นแนวปฏิบัติไม่ได้
ก็เลยคิดว่าจะลองพยายามเขียนอะไรซักอย่างขึ้นมาอุดช่องว่าง
สำหรับคนที่พอเข้าใจการทำงานของฐานข้อมูลอยู่แล้ว แต่แค่ไม่ได้ทำงานด้วย Ruby on Rails มาก่อน
ให้สามารถเห็นภาพการทำงานที่ชัดขึ้น

## Active Record

Model ใน MVC ของ Ruby on Rails นั้นใช้ library ORM เฉพาะของตัวเองที่ชื่อว่า `Active Record`
โดยชื่อของ library ก็ได้รับอิทธิพลโดยตรงมาจากแนวคิด
[`Active record pattern`](https://en.wikipedia.org/wiki/Active_record_pattern)
ซึ่งแนวคิดหลักของ pattern นี้ คือ

- การเข้าถึง table แต่ล่ะ table จะถูก encapsulate อยู่ภายใน class โดยที่แต่ล่ะ instance ของ class ก็คือตัวแทนของ row หรือ record ของข้อมูลใน table และมีวิธีเข้าถึงข้อมูลแต่ล่ะ column ตามชื่อของ column นั้นโดยตรง

  ```ruby
  user = User.new(firstname: "John")
  user.firstname
  # => John
  ```

- เมื่อ instance ของ class ได้รับการ initiate ขึ้นมาใหม่และถูก save ...จะเท่ากับการ insert
  record ใหม่

  ```ruby
  user = User.new(firstname: "John")
  user.save
  # => INSERT INTO "users" (firstname) VALUES ('John', '2024-06-02 12:34:56', '2024-06-02 12:34:56') RETURNING "id"
  ```

- เมื่อ instance ของ class ได้รับการ load จากข้อมูลที่อยู่เดิมและถูก save ...จะเท่ากับการ
  update record ใหม่

  ```ruby
  user = User.where(firstname: "John").first
  user.save
  # => UPDATE "users" SET "phone" = ?, "updated_at" = ? WHERE "users"."id" = ?  [["phone", "12345"], ["updated_at", "2024-06-02 11:34:54.562867"], ["id", 1]]
  ```

- ชื่อของ table กับชื่อของ class จะตั้งตามหลัก
  [Convention over configuration](https://en.wikipedia.org/wiki/Convention_over_configuration)
  ซึ่งสรุปสั้นๆเร็วๆตรงนี้ได้ตามตัวอย่างนี้

  - ชื่อ class = ชื่อ entity ของสิ่งที่เราต้องการใช้ในรูปเอกพจน์
  - ชื่อ table = ชื่อ entity ของสิ่งที่เราต้องการใช้ในรูปพหูพจน์

  <br/>

  > Ex: entity ของ user จะใช้ชื่อ class ว่า User ภายในไฟล์ชื่อ user.rb และอยู่ใน table ที่ชื่อ users

## Active Record Migration

เป็นเครื่องมือที่ทำมาเพื่อ support การทำงานของ Active Record
เพื่อให้นักพัฒนาสามารถทำงานไปพร้อมๆกันได้ โดยที่สามารถตรวจสอบการเปลี่ยนแปลงของ
Database Schema ได้ตลอด

> คำว่า Database Schema ในที่นี้หมายถึง spec ของฐานข้อมูลนั้นๆ เช่น ชื่อ table รวมถึงชื่อ  column ภายใน, Primary Key, Foreign Key, Index ต่างๆ รวมถึง trigger หรือ store procedure ด้วย

### อธิบายตามโครงสร้างไฟล์

```
- config/
  - database.yml
- db/
  - migrate
    - ....
  - schema.rb
  - seeds.rb
```


### ข้อดีและข้อเสียของการทำ Migration แบบนี้

| ข้อดี | ข้อเสีย|
|-------|------|
| สามารถติดตามการเปลี่ยนแปลงได้ง่าย| การจัดการด้วย code ทั้งหมด อาจไม่เหมาะกับองค์กรที่ใช้ DBA แยก |
| เพิ่ม abstraction layer ขึ้นมาเพื่อรองรับการเขียน code ครั้งเดียวให้ใช้ได้หลายฐานข้อมูล เช่น สามารถเปลี่ยนไปใช้ MySQL จาก PostgreSQL ได้ทันที | การ execute คำสั่งโดยตรงอาจมีผลเสียกับ performance ในระบบฐานข้อมูลขนาดใหญ่ |
| รองรับการ rollback ผ่านคำสั่งที่ program ได้      | ความเสี่ยงจากการทำ data loss ระหว่าง migration  |
| การเปลี่ยน schema แต่ล่ะครั้ง สามารถ test ได้     | หลายๆครั้งก็ต้องอาศัยการเปลี่ยน manual ช่วยด้วยอยู่ดี |



### Alternatively ถ้าไม่ใช้ Ruby on Rails แต่อยากได้แบบนี้จะใช้อะไรได้บ้าง


## แนวทางที่ใช้ในการพัฒนา

### คิดด้วย pattern ของ Active Record



### ถ้าจำนวนไฟล์ migrations เยอะขึ้นจนจัดการไม่ไหว

### `schema.rb` คือ Single Source of Truth ของ Database Schema

`rails db:schema:load`

`rails db:schema:dump`
