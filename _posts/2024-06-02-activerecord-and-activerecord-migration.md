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
  user.class.name
  # => "User"
  user.firstname
  # => "John"
  ```

- เมื่อ instance ของ class ได้รับการ initiate ขึ้นมาใหม่และถูก save ...จะเท่ากับการ insert
  record ใหม่

  ```ruby
  user = User.new(firstname: "John")
  user.save
  # => INSERT INTO "users" (firstname) VALUES ("John", "2024-06-02 12:34:56", "2024-06-02 12:34:56") RETURNING "id"
  ```

- เมื่อ instance ของ class ได้รับการ load จากข้อมูลที่อยู่เดิมและถูก save ...จะเท่ากับการ
  update record ใหม่

  ```ruby
  user = User.where(firstname: "John").first
  user.firstname = "Jack"
  user.save
  # => UPDATE "users" SET "firstname" = ?, "updated_at" = ? WHERE "users"."id" = ?  [["firstname", "Jack"], ["updated_at", "2024-06-02 11:34:54.562867"], ["id", 1]]
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

```sh
- config/
  - database.yml
    # ใช้ประกาศค่าต่างๆที่ใช้ในการเชื่อมต่อกับฐานข้อมูล เช่น host, port หรือ
    # connection string
- db/
  - migrate
    # directory ที่ใช้เก็บไฟล์ migration
    # ใช้งานผ่านคำสั่ง rails db:migrate, rails db:rollback
    - ....
  - schema.rb
    # ไฟล์ที่เก็บ schema ปัจจุบันของแอปพลิเคชัน
    # ใช้งานผ่านคำสั่ง rails db:schema:load, rails db:schema:dump
  - seeds.rb
    # ไฟล์ที่ใช้เก็บคำสั่งสำหรับสร้างข้อมูลเบื้องต้นที่จำเป็นในแอปพลิเคชัน
    # ใช้งานผ่านคำสั่ง rails db:seed
```

## การใช้งานในการทำงานจริงๆเป็นยังไง ?

จากที่เกริ่นมาทั้งหมดข้างบน คิดว่าหลายคนสามารถอ่านจนเข้าใจได้ไม่มีปัญหา แต่ถ้าจะมีปัญหาก็คือ
แนวทางปฏิบัติตอนที่ทำงานจริง ว่าเป็นยังไงกันแน่ ทำไมลองทำเองดูแล้วมันก็ยังงงๆ
ไม่แน่ใจว่าจะใช่แบบที่ทำรึเปล่า ซึ่งจากประสบการณ์แล้ว ก็ไม่แปลกที่จะงงกันตรงนี้
เพราะคนที่ชอบใช้ Rails แต่ไม่เข้าใจการทำ migration มีอยู่จำนวนไม่น้อยเลย

> มีทั้งแบบที่แค่หยุดทำไฟล์ migration กับอีกแบบคือหนีไปใช้ NoSQL แบบ MongoDB แทน

### แบบฝึกให้คิดตาม: แอปพลิเคชันซื้อขายสินค้า

สมมติว่า เรากำลังทำแอปพลิเคชันที่สามารถวางขายสินค้าและเปิดให้ผู้ซื้อสามารถสั่งซื้อ
ครั้งละหลายรายการได้ โดยมี entity ที่เราจะใช้ตามนี้

- Buyer ผู้ซื้อสินค้า
- Product สินค้า
- Order รายการสั่งซื้อสินค้า

> ต่อจากนี้ไป ทุกครั้งหลังการสร้าง model, migration ให้รันคำสั่ง
> `rails db:migrate`

### ออกแบบ class ที่จะใช้ ตามลำดับโปรแกรมที่เขียน

> สมมติว่าเราเริ่มจากการมีสินค้าในระบบก่อน

สร้าง model ของ `Product` โดยการ generate model
จะเป็นการสร้างไฟล์ model และ migration ขึ้นมา

```sh
rails generate model product name:string description:text price:integer
```

ก่อนจะไปต่อกันสามารถลองสร้างไฟล์ migration เพื่อเพิ่ม column ง่ายๆได้ เช่น การให้ `Product`
สามารถเก็บจำนวน stock ได้

```sh
rails generate migration add_stock_to_products stock:integer
```

> จะเห็นได้ว่าเราไม่ได้ระบุชื่อ table เป็น parameter
> โดยตรงแต่คำสั่งนี้ก็ยังทำงานตามที่เราต้องการได้ เพราะรูปแบบคำสั่งก็ generate
> จากชื่อไฟล์ตามแนวคิด Convention over configuration
> (จะ generate ไฟล์เปล่าแล้วเขียนเองก็ได้)

มีสินค้าแล้วเราก็อยากให้การสั่งซื้อสินค้าแต่ล่ะครั้งของ `Buyer` ถูกเก็บไว้ใน `Order`

```sh
rails generate model buyer email:string
rails generate model order buyer:references
```

references คือการสร้าง foreign key พร้อมกับ index ไปที่ resource นั้น
ซึ่ง `buyer:references` ในที่นี้ จะทำให้ table `orders`
มี column `buyer_id` เพิ่มเข้าไปพร้อมกับสร้าง index ให้ โดยไฟล์ `order.rb`
ที่ถูกสร้างขึ้นทีหลังจะมีความสัมพันธ์ระบุ `belongs_to :buyer` ระบุให้เลย
แต่ในไฟล์ `buyer.rb` เราต้องเข้าไปใส่ ความสัมพันธ์เองตามข้างล่าง

```ruby
class Buyer < ApplicationRecord
  has_many :orders
end
```

พอมาถึงตรงนี้ เราอยากจะให้ความสัมพันธ์ของ `Order` กับ `Product` เป็นแบบ many-to-many
ซึ่งจากประสบการณ์ของผมค้นพบว่าส่วนมากจะหมดใจกันแถวๆนี้แหละ ที่ไม่รู้จะไปไงต่อดี 😅
จะไปต่อกับกับไฟล์ migration ก็งงๆ จะข้ามส่วนนี้ไปเลยก็เท่ากับว่าไฟล์ก็มีขั้นตอนไม่ครบแล้ว
ฝืนทำต่อไปก็ไม่สมบูรณ์อยู่ดี

แต่เราลองทำต่อตรงนี้ซักหน่อย เพราะ Active Record Migration
มีวิธีใช้งานที่ผมเองเชื่อว่าครอบคลุม 99% ของการใช้งานทั้งหมดแล้ว
กรณีแบบนี้ก็เช่นกัน

### สร้างไฟล์ migration สำหรับ Join Table

เราสามารถสร้างความสัมพันธ์แบบ many-to-many ได้ ด้วยการสร้าง `join table`
ซึ่งใช้คำสั่งสร้างแบบนี้

```sh
rails generate migration CreateJoinTableOrdersProducts order product
```

หลังการรัน `rails db:migrate` ในครั้งนี้ เนื่องจากเราไม่ได้ generate model ขึ้นใหม่ เราจึงจำเป็นต้อง  ประกาศความสัมพันธ์เองภายในไฟล์ model ของ `Product`, `Order`

```ruby
class Product < ApplicationRecord
  has_and_belongs_to_many :orders
end
```

```ruby
class Order < ApplicationRecord
  belongs_to :buyer
  has_and_belongs_to_many :products
end
```

มาถึงตรงนี้เราก็น่าจะสามารถลองเล่นกับความสัมพันธ์ใน `rails console` ดูได้แบบนี้

```ruby
b = Buyer.new(email: 'buyer@test.com')
b.save
# Insert Buyer b

p = Product.new(name: 'The Shirt')
p.save
# Insert Product p

o = Order.new
o.products << p
o.buyer = b
o.save
# Insert Order o with Buyer b, Product p
```

### สิ่งที่อยากให้สนใจหลังแบบฝึกข้างบน

อยากให้ลองศึกษาไฟล์ migration ที่ถูกสร้างขึ้นมาเรียงตามลำดับแล้ว
อาจจะลองรันคำสั่ง `rails db:rollback` เปลี่ยนแปลงค่าบางอย่าง แล้วลอง `rails db:migrate`
อีกครั้งเพื่อดูผลลัพธ์ที่เกิดขึ้น โดยดูตัวอย่างทั้งหมดจาก
[Official Doc](https://guides.rubyonrails.org/active_record_migrations.html)

และศึกษาการเปลี่ยนแปลงของไฟล์ `db/schema.rb` ไปพร้อมๆกัน

### `schema.rb` คือ Single Source of Truth ของ Database Schema

หนึ่งในจุดที่ผมจะเข้าไปเปิดดูเป็นไฟล์แรกๆ หากได้รับมอบหมายให้ทำงานใน Ruby on Rails
แอปพลิเคชัน ก็คือไฟล์ `schema.rb` เพราะเป็นไฟล์ที่เราสามารถมองเห็น
ภาพรวมความซับซ้อนของระบบฐานข้อมูลภายในได้ภายในไฟล์เดียว

### ถ้าจำนวนไฟล์ migrations เยอะขึ้นจนจัดการไม่ไหว

หนึ่งในข้อกังวลอย่างหนึ่งของการเริ่มทำไฟล์ migration ที่ค่อนข้างมีเหตุผล แต่ Active Record
Migration ก็วางแนวทางไว้รับมือแล้ว โดยที่เราสามารถ...

ย้ายไฟล์ที่เก่ามากไปวางที่อื่นหรือลบทิ้งได้เลย แล้วใช้วิธีสร้างฐานข้อมูลเริ่มต้นด้วยคำสั่ง
`rails db:schema:load` ให้สร้างฐานข้อมูลตามข้อมูลใน `schema.rb` แทนการใช้คำสั่ง
`rails db:migrate` เพื่อรันคำสั่งทั้งหมดทีล่ะคำสั่ง

`rails db:schema:dump` คือคำสั่งตรงข้ามของการ load เป็นคำสั่งที่เอาข้อมูล schema ในฐานข้อมูลที่เชื่อมต่อลงเก็บที่ `schema.rb`

### การตรวจสอบสถานะของ database migration

ใช้คำสั่ง `rails  db:migrate:status` เพื่อเรียกดูข้อมูลใน terminal
โดยจะเห็นแบบตัวอย่างข้างล่าง

```sh
 Status   Migration ID    Migration Name
--------------------------------------------------
   up     20240508063759  ********** NO FILE **********
   up     20240602151931  Create products
   up     20240602154948  Create buyers
   up     20240602155002  Create orders
   up     20240602161438  Create join table orders products
  down    20240602161959  Add stock to products
```

### ข้อดีและข้อเสียของการทำ Migration แบบนี้

| ข้อดี | ข้อเสีย|
|-------|------|
| สามารถติดตามการเปลี่ยนแปลงได้ Git Version Control เพราะการเปลี่ยนแปลงทั้งหมดถูกทำผ่าน code| การจัดการด้วย Ruby code ทั้งหมด อาจไม่เหมาะกับองค์กรที่ใช้ DBA แยก เพราะ DBA อาจจะเคืองถ้าโดนบังคับให้เรียน Ruby เพิ่ม 😂|
| เพิ่ม abstraction layer ขึ้นมาเพื่อรองรับการเขียน code ครั้งเดียวให้ใช้ได้หลายฐานข้อมูล เช่น สามารถเปลี่ยนไปใช้ MySQL จาก PostgreSQL ได้ทันที | การ execute คำสั่งโดยตรงอาจมีผลเสียกับ performance ในระบบฐานข้อมูลขนาดใหญ่ ต้องวางแผนให้ดี |
| รองรับการ migrate ละ rollback ผ่านคำสั่งที่ program ได้      | ความเสี่ยงจากการทำ data loss ระหว่าง migration เหมือนข้างบน |
| การเปลี่ยน schema แต่ล่ะครั้ง สามารถ test ได้     |  |


### Alternatively ถ้าไม่ใช้ Ruby on Rails แต่อยากได้แบบนี้จะใช้อะไรได้บ้าง

- Liquibase
- Atlas
- golang-migrate
- Gorm
- Prisma
- etc.
