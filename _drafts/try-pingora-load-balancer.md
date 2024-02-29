---
layout: post
title: ทดลองทำ Load Balancer ด้วย Rust บน Pingora Framework
tags: rust loadbalancer pingora
---

## Pingora Framework
Pingora คือ framework บนภาษา Rust ที่ออกแบบมาเพื่อให้พัฒนา proxy, load balancer ด้วยภาษา rust ได้สะดวก เพื่อรองรับการใช้งานที่ต้องการการปรับแต่งขั้นสูง โดยตอนนี้ยังรองรับเฉพาะการพัฒนาบนภาษา Rust เท่านั้น

## เปรียบเทียบกับ Nginx, Caddy, HAProxy
Pingora ไม่ได้ถูกออกแบบมาเพื่อใช้เป็น web server โดยเฉพาะ จึงไม่มีเครื่องมือหลายอย่างที่มีใน Nginx, Caddy โดยจะใกล้เคียงกับการนำไปใช้แทนที่ HAProxy มากกว่าสองอย่างแรก

## การติดตั้งและการเริ่มต้นใช้งาน
- ทำการสร้าง
  ```shell
  cargo new pingora_lb
  ```

## แหล่งข้อมูลอ้างอิง
- [เหตุผลที่ Cloudflare เลือกสร้าง Pingora](https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet)
- [Github Repository ของ Pingora](https://github.com/cloudflare/pingora/tree/main)

