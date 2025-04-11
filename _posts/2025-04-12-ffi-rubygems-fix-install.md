---
layout: post
title: แก้ปัญหา install ffi gem ไม่ได้ใน MacOS
tags: ruby
date: 2025-04-12 01:53 +0700
---
```sh
brew install libffi

export LDFLAGS="-L$(brew --prefix libffi)/lib"
export CPPFLAGS="-I$(brew --prefix libffi)/include"
```
