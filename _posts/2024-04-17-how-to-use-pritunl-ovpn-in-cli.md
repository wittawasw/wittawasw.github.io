---
layout: post
title: วิธีใช้ ovpn profile ที่มี pin code ของ Pritunl ด้วย openvpn ผ่าน command-line
tags: pritunl, vpn, openvpn, cli
description: นั่งหาวิธีการใช้ VPN profile ของ Pritunl ด้วย CLI เพราะเบื่อการที่ Pritunl Client มักจะหลุดและต้องมาคอยใส่ pin ใหม่อยู่เรื่อยๆ
date: 2024-04-17 16:47 +0700
---

นั่งหาวิธีการใช้ VPN profile ของ Pritunl ด้วย CLI เพราะเบื่อการที่ Pritunl Client มักจะหลุดและต้องมาคอยใส่ pin ใหม่อยู่เรื่อยๆ

### คำสั่งในการรัน VPN

ใช้ client ของ openvpn โดยถ้าใน MacOS สามารถ install ผ่าน [Homebrew](https://formulae.brew.sh/formula/openvpn) ได้

```sh
# เมื่อ profile ที่ใช้อยู่ใน file ชื่อ profile.vpn
# และ username, pin code อยู่ใน pass.txt
# --daemon คือ การสั่งให้ daemonize process
sudo openvpn --config profile.ovpn --auth-user-pass "pass.txt" --daemon
```

### ตัวอย่างไฟล์ `pass.txt`

- เมื่อ username = username@example.com
- และ pin code = 123456

```sh
username@example.com
123456
```
