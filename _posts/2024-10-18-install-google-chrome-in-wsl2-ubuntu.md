---
layout: post
title: Install Google Chrome ใน wsl2 Ubuntu
tags: chrome cli ubuntu wsl2
keywords: chrome, cli, ubuntu, wsl2
description: ''
date: 2024-10-18 10:34 +0700
---

เจอ error นี้ระหว่างที่กำลังจะใช้ CLI ในการรันคำสั่งที่ต้องเรียก Chrome ออกมาใช้
แล้วก็ได้รู้ว่า wsl2 มันยังไม่ฉลาดพอที่จะเรียกใช้ Chrome จาก Windows ให้เอง
ซึ่งก็ไม่แปลกอะไร ก็เลยได้ลอง install Chrome ให้ใช้ได้ด้วยโปรแกรมที่เป็น CLI ดู

```
[  WARN ] The CHROME_PATH environment variable must be set to a Chrome/Chromium executable no older than Chrome stable.
[ ERROR ] You have to install Google Chrome, Chromium, or Microsoft Edge.
```

ลองพยายามแบบสั้นๆ เร็วๆดู

```sh
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt -y install ./google-chrome-stable_current_amd64.deb
```

เจอ error

```
N: Download is performed unsandboxed as root as file '/home/wittawasw/honeypot/wittawasw.github.io/google-chrome-stable_current_amd64.deb' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
```

แปล error ออกมาได้ว่า path ที่เราใช้อยู่ ไม่สามารถเข้าถึงได้ด้วย user _apt
เพราะฉะนั้นเราก็แค่ต้องย้ายไฟล์ .deb ไป execute ที่ไหนก็ได้ที่ _apt สามารถเข้าถึงได้
จะสร้างใหม่แล้ว chown ก็ได้ แต่ที่ /tmp ก็เป็นที่ที่ดีที่จะเอาของวางชั่วคราวแล้วก็มีอยู่แล้ว
ไม่ต้องสร้างใหม่

```sh
mv google-chrome-stable_current_amd64.deb /tmp
cd /tmp
sudo apt -y install ./google-chrome-stable_current_amd64.deb

# ลบไฟล์ทิ้งแล้วตรวจดูว่าที่ install ไป คือ chrome version อะไร
rm google-chrome-stable_current_amd64.deb
google-chrome --version
```

เรียกใช้งานได้ แต่ให้แสดงผลไม่ได้ จะติด error

```
[648041:648140:1018/101730.158350:ERROR:object_proxy.cc(576)] Failed to call method: org.freedesktop.UPower.EnumerateDevices: object_path= /org/freedesktop/UPower: org.freedesktop.DBus.Error.ServiceUnknown: The name org.freedesktop.UPower was not provided by any .service files
glx: failed to create drisw screen
```

ถ้าพยายามไปต่อจากนี้ติดปัญหาหลายอย่างที่คิดว่า ยังไม่คุ้มที่จะทำ บวกกับยังไม่ต้องการดูผลลัพธ์ของ Chrome แต่แค่ต้องการเอามารัน CLI ก็เลยปล่อยไปแค่นี้
