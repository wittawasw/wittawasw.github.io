---
layout: post
title: ทดลองทำ Load Balancer ด้วย Rust บน Pingora Framework
tags: rust loadbalancer pingora
---

> Code ตัวอย่างของบทความนี้ คัดลอกมาจาก [quickstart guide](https://github.com/cloudflare/pingora/blob/main/docs/quick_start.md) ของ Pingora

## Pingora Framework
Pingora คือ framework บนภาษา Rust ที่ออกแบบมาเพื่อให้พัฒนา proxy, load balancer ด้วยภาษา rust ได้สะดวก เพื่อรองรับการใช้งานที่ต้องการการปรับแต่งขั้นสูง โดยตอนนี้ยังรองรับเฉพาะการพัฒนาบนภาษา Rust เท่านั้น

### เปรียบเทียบกับ Nginx, Caddy, HAProxy
Pingora ไม่ได้ถูกออกแบบมาเพื่อใช้เป็น web server, proxy เพื่อให้พร้อมใช้งานเหมือนกับทั้งสามอย่าง แต่เป็น framework สำหรับนักพัฒนาโดยเฉพาะ จึงไม่สามารถเทียบกันโดยตรงได้ แต่อนาคตอันใกล้ก็น่าจะมีคนเอามาสร้างเป็น web server, proxy ที่พร้อมใช้งาน

> จากตรงนี้เป็นต้นไป ผู้อ่านควรมีความรู้เบื้องต้นเกี่ยวกับ Network และเครื่องมือของภาษา Rust

## การติดตั้งและการเริ่มต้นใช้งาน
-
- ทำการสร้าง
  ```shell
  cargo new pingora_lb
  ```

## แหล่งข้อมูลอ้างอิง
- [เหตุผลที่ Cloudflare เลือกสร้าง Pingora](https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet)
- [Github Repository ของ Pingora](https://github.com/cloudflare/pingora/tree/main)

