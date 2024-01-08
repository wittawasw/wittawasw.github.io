---
layout: post
title: 'ลองใช้ Hotwired: Turbo Drive ในเว็บที่ไม่ใช่ Rails'
tags: rails, turbo, hotwired, html, javascript
date: 2024-01-08 12:13 +0700
---
![hotwired](/images/post/hotwired.png)

[Hotwired](https://hotwired.dev/) คือชุด Javascript library ที่เป็นค่าเริ่มต้นของ Ruby on Rails ตั้งแต่เวอร์ชั่นที่ 7 เป็นต้นมา
แต่การที่ library ถูกเขียนด้วย Javascript ทั้งหมด การใช้งานจึงไม่ได้ถูกผูกอยู่กับ Ruby on Rails เท่านั้น
โดย [Turbo Drive](https://turbo.hotwired.dev/handbook/drive) เป็น feature แรกของ Hotwired ที่เริ่มใช้งานได้ง่ายที่สุด

## การติดตั้งในหน้าเว็บเพื่อเริ่มใช้งาน

- วาง script tag นี้ที่หน้า HTML layout ของทุกหน้า
  ```html
  <script type="module">
    import * as Turbo from 'https://cdn.skypack.dev/@hotwired/turbo@7.3.0';
  </script>
  ```
- ตัวอย่าง
  ```html
  <html>
    <head>
      <title>...</title>
      .
      .
      .
      <script type="module">
        import * as Turbo from 'https://cdn.skypack.dev/@hotwired/turbo@7.3.0';
      </script>
    </head>
    <body>
      .
      .
      .
    </body>
  </html>
  ```
- เท่านี้ก็สามารถเริ่มใช้งาน Turbo Drive ได้เลย

> ตั้งค่าเวอร์ชั่นของ turbo ที่ 7.3.0 ซึ่งเป็นรุ่น stable

## วิธีใช้งาน
- ไม่ต้องเพิ่มคำสั่งอะไรใน source code เพราะการ import turbo ด้วย script tag ก็ทำให้ Turbo Drive เริ่มทำงานแล้ว

> แล้วต่างจากเดิมยังไง ในเมื่อยังไม่ได้ทำอะไรเพิ่ม ?

## สิ่งที่เปลี่ยนไปหลังจากเพิ่ม script
- การ navigation ผ่าน tag `<a>` และ `<form>` ทุกอันจะถูกแทนที่ด้วย Turbo Drive
- การ navigate ด้วย Turbo Drive จะทำงานคล้ายกับ SPA (Single Page Application)
  แต่แตกต่างตรงที่เราไม่ต้องเขียน logic เอง และ Turbo Drive จะทำการ replace HTML ภายใน
  `<body>` ให้โดยอัตโนมัติ โดยไม่มีการ refresh page
- หากทำการ inspect ที่ network tab จะเห็นความแตกต่าง ตามตัวอย่างข้างล่าง

<div style="display: flex; justify-content: space-between;margin-bottom:50px;">
  <div>
    <video width="98%" preload="auto" muted controls>
      <source src="/videos/normal-navigation.mp4" type="video/mp4">
      Your browser does not support the video tag.
    </video>
    <quote>
      navigation แบบปกติก่อนใส่ script
    </quote>
  </div>
  <div>
    <video width="98%" preload="auto" muted controls>
      <source src="/videos/navigation-with-turbodrive.mp4" type="video/mp4">
      Your browser does not support the video tag.
    </video>
    <quote>
      navigation ด้วย Turbo Drive หลังใส่ script
    </quote>
  </div>
</div>

> สามารถทำการ inspect network และ script tag ของเว็บนี้เพื่อดูเป็นตัวอย่างได้
