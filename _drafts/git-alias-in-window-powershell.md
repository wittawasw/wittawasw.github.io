---
layout: post
title: ตั้งค่าตัวย่อของ Git CLI บน Windows Powershell
tags: git cli windows, powershell
keywords: git, cli, windows, powershell
description:
---

หลังจากที่มีเหตุให้ต้องใช้ Windows เป็นเครื่องสำรองอยู่เกือบเดือน
หลายๆอย่างที่ใช้งานใน Ubuntu ก็ย้ายมาอยู่บน WSL 2 ได้เยอะแล้ว
ช่วงนี้มีความคิดอยากทำ Flutter app ก็เลยอยากลงบนเครื่อง Windows ด้วย
แต่ก็ติดที่มันใช้ผ่าน WSL ไม่สะดวก ถ้าไม่ผ่าน WSL ก็ติดเรื่องเดียว คือ
ไม่รู้ว่าจะหาทางเอา Git CLI แบบตัวย่อที่ปกติใช้งานบน Ubuntu มาทำบน Windows
ยังไง รู้แต่ว่ายังไงก็ไม่ยอมใช้ GUI

ก็เลยได้ลองเล่น Powershell ไปมาก็รู้สึกลงตัวกับแบบนี้
ยังไม่รู้สึกดีเท่าที่ทำงานบน MacOS, Ubuntu แต่ก็พอใช้ได้แล้ว

- ใช้งาน terminal ใน vscode โดยการเรียกผ่าน `Ctrl + ~`
- ตั้งค่า Alias กับ Run Flutter command ผ่าน terminal เลย


## ไฟล์ที่ใช้ตั้งค่าให้ powershell

```sh
# เปิดไฟล์ที่ $profile ด้วย VS Code
code $profile
```

## Alias ที่ใช้

> ลอกมาจาก `Oh-my-zsh`

```sh
# ประกาศด้วยคำสั่ง Function
Function gitCommand { git $args }
Function gitCommit { git commit --verbose }
Function gitCheckout { git checkout $args }
Function gitCheckoutBranch { git checkout -b $args }
Function gitRemote { git remote -v }
Function gitRemoteRemove { git remote remove $args }
Function gitRemoteAdd { git remote add $args }
Function gitPull { git pull origin $args }

Function gitPush {
    param (
        [string[]]$args
    )

    if ($args.Count -ne 0 -and $args.Count -ne 1) {
        git push origin $args
    } else {
        if ($args.Count -eq 0) {
            $b = (git rev-parse --abbrev-ref HEAD)
        }
        git push origin $($b = if ($args.Count -eq 0) { $b } else { $args[0] })
    }
}

Function gitPushForce {
    param (
        [string]$arg
    )

    if ($args.Count -ne 1) {
        $b = (git rev-parse --abbrev-ref HEAD)
    }
    git push --force origin $b
}

Function gitFetch { git fetch --all --tags --prune --jobs=10 }
Function gitStatus { git status }
Function gitSwitch { git switch $args }
Function gitResetHard { git reset --hard $args }
Function gitResetHardHead { git reset --hard HEAD }
Function gitBranch { git branch $args }
Function gitBranchDelete { git branch -d $args }
Function gitLogGraph { git log --oneline --graph --decorate }
Function gitLogAllGraph { git log --oneline --all --graph --decorate }
Function gitAdd { git add $args }
Function gitAddAll { git add . }

# ตั้งค่า alias ให้ชี้ไปที่ Function
Set-Alias -Name g -Value gitCommand

# ใช้ gc ไม่ได้เพราะทับกับคำสั่งเดิมของ windows เลยต้องใช้เป็น gcc แทน
Set-Alias -Name gcc -Value gitCommit
Set-Alias -Name gco -Value gitCheckout
Set-Alias -Name gcb -Value gitCheckoutBranch
Set-Alias -Name grv -Value gitRemote
Set-Alias -Name grrm -Value gitRemoteRemove
Set-Alias -Name gra -Value gitRemoteAdd
Set-Alias -Name ggl -Value gitPull
Set-Alias -Name ggp -Value gitPush
Set-Alias -Name ggf -Value gitPushForce
Set-Alias -Name gst -Value gitStatus
Set-Alias -Name gsw -Value gitSwitch
Set-Alias -Name grh -Value gitResetHard
Set-Alias -Name grhh -Value gitResetHardHead
Set-Alias -Name gb -Value gitBranch
Set-Alias -Name gbd -Value gitBranchDelete
Set-Alias -Name glgg -Value gitLogGraph
Set-Alias -Name glola -Value gitLogAllGraph
Set-Alias -Name gf -Value gitFetch
Set-Alias -Name ga -Value gitAdd
Set-Alias -Name gaa -Value gitAddAll

```

## การLoad ใส่ powershell

```sh
# กดบันทึกแล้วเปิดใหม่หรือใช้คำสั่งนี้ เป็นเหมือนการสั่ง source ใน Unix
. $profile

# ตั้งค่าให้ ggp, ggf สามารถทำงานได้โดยไม่ต้อง set upstream
g config --global push.autoSetupRemote true
```
