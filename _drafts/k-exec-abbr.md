---
layout: post
title: k-exec-abbr
tags: kubernetes cli
keywords: kubernetes, cli
description: ''
header-img: "/images/post/ar-ar-migrate.png"
---


```sh
function ke() {
  POD_NAME=$(kgpo | grep "$1" | awk '{print $1}')
  if [ -z "$2" ]; then
    k exec -it "$POD_NAME" -- /bin/sh
  else
    shift
    k exec -it "$POD_NAME" -- "$@"
  fi
}
```

```sh
function kl() {
  POD_NAME=$(kgpo | grep "$1" | awk '{print $1}')
  if [ "$2" = "-f" ]; then
    k logs -f "$POD_NAME"
  else
    k logs "$POD_NAME"
  fi
}
```

```sh
function krrm() {
  POD_NAMES=$(kgpo | grep "$1" | awk '{print $1}')

  for POD_NAME in $POD_NAMES; do
    krm pod "$POD_NAME"
    echo "Pod $POD_NAME restarted."
  done
}
```
