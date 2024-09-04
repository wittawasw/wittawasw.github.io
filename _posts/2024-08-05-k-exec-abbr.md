---
layout: post
title: การใช้ kubectl แบบ "ตัวย่อ"
tags: kubernetes cli
keywords: kubernetes, cli
description: 'เนื้อหานี้เป็นส่วนหนึ่งของที่ผมใช้ทำงานจริง แล้วก็ได้เอาไปแจกเป็นโพยให้ที่ทำงานเดิมมาแล้ว
  เอา copy มาลง blog ส่วนตัวเพื่อเก็บไว้เป็นความทรงจำ '
date: 2024-08-05 21:52 +0700
---
เนื้อหานี้เป็นส่วนหนึ่งของที่ผมใช้ทำงานจริง แล้วก็ได้เอาไปแจกเป็นโพยให้ที่ทำงานเดิมมาแล้ว
เอา copy มาลง blog ส่วนตัวเพื่อเก็บไว้เป็นความทรงจำ

## Prerequisites

- ความเข้าใจเบื้องต้นเกี่ยวกับ kubectl, k8s
- install kubectl ไว้ในเครื่อง
- install
  - [kubectl-aliases](https://github.com/ahmetb/kubectl-aliases)
  - [kubectx](https://github.com/ahmetb/kubectx) (kubens จะถูก install มาด้วยกัน)


## ข้อดีของการใช้ตัวย่อแบบนี้
ในการใช้งาน kubectl แบบปกติ คำสั่งจะค่อนข้างยาว เช่น

```sh
-> kubectl get pod --context=dev --namespace=authen
```

# หรือ
```sh
-> kubectl config use-context dev
-> kubectl get pod --namespace=authen
```

ซึ่งยาว... ยิ่งคิดว่าอันนี้เป็นคำสั่งที่สั้นที่สุดแล้วยังรู้สึกไม่อยากพิมพ์ ถ้าแบบที่ซับซ้อนกว่านี้ ยิ่งไม่อยากพิมพ์
โดยจากข้างบน ถ้าเราลงเครื่องตามข้างบนจะเหลือแบบนี้

```sh
# ทำการ set context, namespace แค่ครั้งเดียว
-> k ctx dev
-> k ns authen

# จากนั้นใช้คำสั่งต่อเนื่องได้เลย ไม่ต้องพิมพ์ namespace อีก
-> kgpo
```

สั้นกว่ากันเยอะ...  เวลาที่ใช้พิมพ์ลดลงได้หลายวินาที
โดยรายละเอียดทั้งหมดสามารถดูคำสั่งได้จาก source ใน [github](https://github.com/ahmetb/kubectl-aliases/blob/master/.kubectl_aliases)

## ข้อเสีย
จำคำสั่งจริงไม่ได้


## Use case ใช้งานจริง โดยเอาไปต่อยอดด้วย shell script

การค้นหา pod
หลายๆครั้ง เราอยาก ดูสถานะหรือดูจำนวน pod หรืออาจจะแค่อยากดูชื่อเต็มของ pod เราสามารถ get pod ออกมาทั้งหมด แล้ว grep

```sh
# ดู pod ทั้งหมด
-> kgpo
# NAME                                         READY   STATUS    RESTARTS        AGE
# auth-service-api-657b4c4658-cs76p            1/1     Running   0               77m
# core-service-5498d9f57c-2tll1                1/1     Running   0               20d
# core-service-5498d9f57c-2tll2                1/1     Pending   0
# report-service-8745c86f5-w24dd               1/1     Running   0               13h
# shared-service-6f88d74cdc-thpfw              1/1     Running   3 (4h12m ago)   12h

# ดู pod ทั้งหมด ที่มีคำว่า core
-> kgpo | grep core
# core-service-5498d9f57c-2tll1                1/1     Running   0               20d
# core-service-5498d9f57c-2tll2                1/1     Pending   0

# ดู pod ทั้งหมด ที่มีคำว่า core และไม่มีคำว่า Running
# เอาไว้หา pod ที่ผิดปกติ
-> kgpo | grep core | grep -v Running
# core-service-5498d9f57c-2tll2                1/1     Pending   0
```

ถ้าอยากได้แค่ชื่ออย่างเดียว ก็ต่อยอดด้วยคำสั่ง awk แล้วทำเป็นฟังก์ชันไปใส่ใน shell

```sh
function ks() {
  kgpo | awk 'NR>1 {print $1}' | grep "$1" | while read -r POD_NAME; do
    echo "$POD_NAME"
  done
}

-> ks core
# core-service-5498d9f57c-2tll1
# core-service-5498d9f57c-2tll2
```

ต่อยอดจากการค้นหา pod โดยสั่งให้ ลบ pod ที่เราหาชื่อเจอออก
การสั่งลบ Pod ตามชื่อนั้นๆ เพื่อให้ Deployment ทำการสร้าง Pod ใหม่มาแทนที่ เพราะฉะนั้นก็เท่ากับเป็นการ restart application process แบบเร็วไปในตัว

```sh
function krrm() {
  kgpo | awk 'NR>1 {print $1}' | grep "$1" | while read -r POD_NAME; do
    # คำสั่ง krm ย่อมาจาก kubectl delete
    krm pod "$POD_NAME"
    echo "pod $POD_NAME restarted."
  done
}

# ตัวอย่างการใช้
-> krrm core
# pod core-service-5498d9f57c-2tll1 deleted
# pod core-service-5498d9f57c-2tll1 restarted.
# pod core-service-5498d9f57c-2tll2 deleted
# pod core-service-5498d9f57c-2tll2 restarted.
```

## การ exec จากข้างใน pod
ในกรณีที่เราต้องการ exec เข้าไปรันคำสั่งบางอย่างข้างใน pod สามารถใช้ฟังก์ชั่น


```sh
# ke โดย default จะสั่ง /bin/sh เพื่อเข้าไปข้างใน แต่สามารถใส่คำสั่งอื่นต่อท้าย
# เพื่อ override ได้
function ke() {
  POD_NAME=$(kgpo | awk 'NR>1 {print $1}' | grep "$1" | head -n 1)

  if [ -z "$2" ]; then
    k exec -it "$POD_NAME" -- /bin/sh
  else
    shift
    k exec -it "$POD_NAME" -- "$@"
  fi
}

# ตัวอย่างการใช้เพื่อ exec เข้าไปใน pod
# ถ้าอยากเข้าให้ตรง pod แบบเจาะจงก็เอาชื่อเต็มๆมาใส่
-> ke core-service-5498d9f57c-2tll1
# /root/app/ |

# ถ้าแค่อยากเข้าอันไหนก็ได้เพราะมันควรจะเหมือนกันทุกอันก็พิมพ์ชื่อแค่ส่วนเดียวได้
-> ke core
# /root/app/ |
```


ตัวอย่าง use case จริง ของการ exec จากข้างใน pod
ถ้าเราต้องการตรวจสอบว่า Environment Variables ที่เรา deploy ที่ pod มีค่าถูกต้องหรือไม่ สามารถใช้คำสั่ง ke ข้างบนแบบนี้

```sh
# เรียกดู env vars ทั้งหมดใน pod
-> ke core env

# เรียกดู env vars ใน pod ที่นำหน้าว่า DB_
-> ke core env | grep DB_
```

ในการทำงานจริงมักจะใช้คู่กับการ restart pod ข้างบน โดย

- ตรวจดูว่า env ใน pod แล้วพบว่าผิด หรือยังไม่มี env
- deploy Env ใหม่ ผ่าน gitops
- ลบ pod เพื่อ restart deployment
- ตรวจดูว่า env ใหม่ ได้ถูก deploy ถูกต้องแล้วหรือไม่

## การดู log ของ pod
วิธีประยุกต์ใช้ script คล้ายกับการ exec แต่ลองเขียนใหม่สำหรับดู log

```sh
function kl() {
  POD_NAME=$(kgpo | awk 'NR>1 {print $1}' | grep "$1" | head -n 1)

  if [ "$2" = "-f" ]; then
    k logs -f "$POD_NAME"
  else
    k logs "$POD_NAME"
  fi
}

# ตัวอย่างการใช้
# เจาะจงชื่อ
-> kl core-service-5498d9f57c-2tll1
# เอาแค่อันแรก
-> kl core

# ใส่ -f เพื่อสั่งให้ process รอรับ log แบบ live
-> kl core-service-5498d9f57c-2tll1 -f
```
