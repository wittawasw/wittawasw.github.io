---
layout: post
title: การใช้ docker สำหรับเตรียมพื้นที่พัฒนาเว็บแอปพลิเคชัน
date: 2022-04-25 00:32 +0700
---
การพัฒนาเว็บแอปพลิเคชันในปัจจุบัน มีความซับซ้อนและความต้องการทางเทคนิคที่เพิ่มขึ้นมาก
การพัฒนาจึงกลายเป็นการทำงานที่ต้องทำร่วมกันกับคนจำนวนมากขึ้น ไม่ได้จำกัดแค่คนกลุ่มเล็กเท่านั้น
แต่อาจหมายถึงการรวมกลุ่มทำงานในทีมที่อาจมีขนาดถึง 50-100 คน
การเตรียมชุดซอฟต์แวร์เพื่อให้ทีมงานจำนวนมากสามารถใช้ร่วมกันได้ โดยที่มีความใกล้เคียงกับพื้นที่จริงบน
production environment จึงกลายเป็นความจำเป็น ซึ่ง docker สามารถช่วยตอบโจทย์ความต้องการนี้ได้
ทว่าการใช้งาน docker จำเป็นต้องเข้าใจพื้นฐานการเชื่อมต่อ network เพื่อให้สามารถตั้งค่าการใช้งานภายในโรแกรมได้เอง ตามที่กำลังจะอธิบายต่อไปด้านล่าง

## การเชื่อมต่อผ่าน docker network
- หากไม่มีการตั้งค่าเพิ่มเติม docker จะทำงานภายในเครือข่าย network ในรูปแบบของ `bridge mode`
  network ที่ใช้ชื่อว่า bridge
  ```shell
  # คำสั่ง cli ที่ใช้ในการดูรายละเอียดของ docker network

  # ดูรายชื่อ docker network
  docker network ls

  # ดูรายละเอียดของ docker network
  docker network inspect <ชื่อของ network>

  # ดูรายละเอียดของ docker network ค่าเริ่มต้น
  docker network inspect bridge
  ```
- การเชื่อมต่อผ่าน docker network ควรสร้าง network ชื่อเฉพาะของแต่ล่ะงานขึ้นมา
  เพื่อแยกระบบออกจากการกันให้ชัดเจน
  ```shell
  # ตัวอย่างการสร้าง network

  # สร้าง network ชื่อ myweb
  ➜  ~ docker network create myweb
  94e2714b100d6634299bbbc7dfb30bf7bc7d150e27f4f48764bec91d32d4686e


  # ทดลองสร้าง container ขื่อ redis-store สำหรับ redis ขึ้นมาบน network ชื่อ myweb
  ➜  ~ docker run --name redis-store --network myweb redis:6-alpine
  1:C # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
  1:C # Redis version=6.2.6, bits=64, commit=00000000, modified=0, pid=1, just started
  1:C # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
  1:M * monotonic clock: POSIX clock_gettime
  1:M * Running mode=standalone, port=6379.
  1:M # Server initialized
  1:M * Ready to accept connections

  ➜  ~ docker ps
  CONTAINER ID   IMAGE            COMMAND                  CREATED              STATUS              PORTS      NAMES
  cb67ff6bd900   redis:6-alpine   "docker-entrypoint.s…"   About a minute ago   Up About a minute   6379/tcp   redis-store


  # ดูรายละเอียดของ docker network ค่าเริ่มต้น
  ➜  ~ docker network inspect myweb
  [
      {
          "Name": "myweb",
          .
          # ของจริงมีรายละเอียดเยอะกว่านี้แต่ย่อมาให้เห็นเฉพาะส่วนของ containers
          .
          "Containers": {
              "cb67ff6bd90061869639e950e908b169d830e54173d479c9ff2186597d2fd1b7": {
                  "Name": "redis-store",
                  "EndpointID": "ed4aa0f9da01cbb7bfb7f6ae6a8eea8af8b58b5c66f16aef9b13831dc652fc22",
                  "MacAddress": "02:42:ac:1b:00:02",
                  "IPv4Address": "172.27.0.2/16",
                  "IPv6Address": ""
              }
          }
      }
  ]

  ```
- จากตัวอย่างด้านบนจะเห็นได้ว่าเราสามารถเข้าถึง container ที่ชื่อว่า redis-store ได้ด้วย
  IP address `172.27.0.2` แต่ทว่า IP address นี้จะเปลี่ยนไปทุกครั้งที่มีการ restart container
  ดังนั้นการนำค่า IP Address ไปใช้งานโดยตรงจึงไม่ใช่เรื่องที่ดีนักเพราะเราต้องเสียเวลามาเปลี่ยน IP Address
  อยู่เรื่อยๆ
- วิธีการที่ถูกต้องคือควรจะนำแอปพลิเคชันที่เรียกใช้ redis ใส่ไว้ใน network ที่ชื่อ myweb ด้วย
  และภายใน code ที่ทำการเชื่อมต่อก็ใช้ชื่อของ container แทน hostname, IP address ได้เลย
  ในตัวอย่างนี้ code ที่ทำการเชื่อต่อ redis จะสามารถเชื่อมต่อกับ redis ผ่านตัวอย่าง URL ต่อไปนี้ได้
  - redis://172.27.0.2:6379  *(ใช้ IP Address โดยตรงก็ได้แต่เป็นวิธีที่ไม่แนะนำ)*
  - redis://redis-store:6379  *(ควรใช้ชื่อ container แทน hostname)*

## การเข้าถึง container ด้วย port mapping
- ในกรณีที่เราต้องการเข้าถึงผ่าน `localhost` หรือ `127.0.0.1` เราสามารใช้ใช้ option -p เพื่อทำการ
  port mapping ออกมาจาก docker network
  ```shell
  # แอปพลิเคชันที่อยู่ภายใน localhost จะสามารถเชื่อต่อได้ผ่าน redis://localhost:6379
  docker run --name redis-store --network myweb -p 127.0.0.1:6379:6379/tcp redis:6-alpine

  # /tcp อนุโลมให้ไม่ใส่ได้เพราะค่าเริ่มต้นคือ tcp อยู่แล้ว และอีก protocol ที่ชื่อ udp ก็ไม่เป็นที่นิยมใช้
  docker run --name redis-store --network myweb -p 127.0.0.1:6379:6379 redis:6-alpine
  ```
- จะสังเกตได้ว่าเราทำการ map ให้เข้ากับ host `127.0.0.1`
  ซึ่งข้อผิดพลาดหรือความไม่ระวังอย่างหนึ่งที่พบได้บ่อยคือการใช้ option -p โดยไม่ระบุ host
  ทำให้ container สามารถถูกเข้าถึงได้จาก host ภายนอกทั้งหมด
  ```shell
  # สังเกตตรง -p 6379:6379 ซึ่งเป็นลักษณะที่ไม่พึงประสงค์ ควรใช้ก็ต่อเมื่อเข้าใจความเสี่ยงเท่านั้น
  docker run --name redis-store --network myweb -p 6379:6379 redis:6-alpine
  ```
-

## บันทึกและติดตามการเปลี่ยนแปลงการตั้งค่า docker ด้วย docker-compose
- การทำเว็บแอปพลิเคชัน องค์ประกอบจะเพิ่มความซับซ้อนมากขึ้นเรื่อยๆ ตามความต้องการของแอปพลิเชัน
  ซึ่งทำให้การตั้งค่า docker ก็ต้องมีการบันทึกเก็บไว้ใน version control
  เพื่อให้สามารถติดตามการเปลี่ยนแปลงได้
- สามารถตั้งค่า docker network จากภายในไฟล์ `docker-compose.yml` ได้เลย
  โดยไม่จำเป็นต้องสร้าง network ด้วยคำสั่ง docker network create <ชื่อ network> เอาไว้ก่อนก็ได้
- docker-compose ไม่ได้เป็นเพียงรูปแบบที่ช่วยบันทึก แต่เรายังสามารถเรียกใช้งานทุกองค์ประกอบได้โดยที่ไม่ต้อง
  docker run ทีล่ะ container ลดความยุ่งยากในการจำตัวเลือกการตั้งค่าต่างๆเหลือเพียงให้จำเพียงชื่อ
  service ของ container เท่านั้น

  ```yaml
  # docker-compose.yml
  version: '3'
  services:
    web:
      build:
        context: .
        dockerfile: ./Dockerfile
      ports:
        - 8080:8080 # port mapping เพื่อให้สามารถใช้ localhost:8080 แทน internal IP ได้
      volumes:
        - .:/var/www/html
      networks:
        - myweb
    redis:
      image: 'redis:6-alpine'
      # ใช้สัญลักษณ์ # ในการ comment เพื่อให้ docker-compose ไม่อ่านค่า
      # comment port mapping ออก โดยสามารถนำเครื่องหมาย # ข้างหน้าออกเพื่อเปิดใช้งานได้ในกรณี
      # ที่ต้องการ debug
      # ports:
      #  - 6379:6379
      environment:
        - ALLOW_EMPTY_PASSWORD=yes
      networks:
        - myweb

  networks:
    myweb:
      driver: bridge
  ```
- การเรียกใช้งานเบื้องต้น
  ```shell
  # เรียกใช้งานในรูปแบบ active process สามารถ kill โดยการกด ctrl-C หรือ ปิด terminal ได้
  docker-compose up

  # เรียกใช้งานในรูปแบบ daemon process
  docker-compose up -d

  # ตัวอย่างการเรียกใช้งานเฉพาะบาง service
  docker-compose up web
  docker-compose up redis

  # ตัวอย่างการเรียกใช้งานหลาย service แบบเจาะจง
  docker-compose up web redis
  ```

## Environment Variables
คำสั่ง docker-compose จะอ่านค่าจากไฟล์ `.env` โดยอัตโนมัติ ทำให้ไม่จำเป็นต้องใส่ค่าบางค่า ที่ผู้พัฒนาบางคนอาจใช้ไม่ตรงกันได้ เช่น port ของแอปพลิเคชัน หรือ ชื่อของ network ที่ใช้

```.env
# .env
EXPOSED_WEB_PORT=8080
WEB_DIRECTORY=.
EXPOSED_REDIS_PORT=6379
MY_WEB_NETWORK_NAME=myweb
```

```yaml
# docker-compose.yml

  version: '3'
  services:
    web:
      build:
        context: .
        dockerfile: ./Dockerfile
      ports:
        - ${EXPOSED_WEB_PORT}:8080
      volumes:
        - ${WEB_DIRECTORY}:/var/www/html
      networks:
        - myweb
    redis:
      image: 'redis:6-alpine'
      # ports:
      #  - ${EXPOSED_REDIS_PORT}:6379
      environment:
        - ALLOW_EMPTY_PASSWORD=yes
      networks:
        - myweb

  networks:
    myweb:
      name: ${MY_WEB_NETWORK_NAME}
      driver: bridge
  ```
