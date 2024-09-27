---
layout: post
title: ใช้ timedatectl ตั้งค่า datetime ในเครื่อง Ubuntu
tags: unix cli ubuntu timedatectl
keywords: unix, cli, ubuntu, timedatectl
description: ''
---

```sh
# ดูเวลาปัจจุบันที่ตั้งไว้
timedatectl
#                Local time: Fri 2024-09-27 12:13:31 +07
#            Universal time: Fri 2024-09-27 05:13:31 UTC
#                  RTC time: Fri 2024-09-27 05:13:31
#                 Time zone: Asia/Bangkok (+07, +0700)
# System clock synchronized: yes
#               NTP service: active
#           RTC in local TZ: no

# ตั้งให้ sync อัตโนมัติผ่าน internet
timedatectl set-ntp yes
timedatectl set-ntp no

# date - YYYY-MM-DD
timedatectl set-time 2024-03-31
# time - HH:MM:SS
timedatectl set-time 12:11:12

# การตั้ง timezone
timedatectl list-timezones
timedatectl set-timezone "<ชื่อ timezone ที่ได้มา>"
timedatectl set-timezone UTC
```
