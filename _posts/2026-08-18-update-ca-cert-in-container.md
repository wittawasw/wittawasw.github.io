---
layout: post
title: Update CA cert in container
tags: docker
date: 2026-08-18 13:14 +0700
---
Some internal network require custom root CA if the proxy is using self-signed.

## Updating CA Cert
```sh
#!/usr/bin/env bash
# set -e

for f in *.cer; do
[ -e "$f" ] || continue
cp -- "$f" "${f%.cer}.crt"
done

sudo cp ./*.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

## For Docker
```docker
COPY tmp/certificates/ /tmp/certificates/
RUN if [ "$CA" = "true" ]; then \
      apk add --no-cache ca-certificates && \
      find /tmp/certificates -type f -name '*.crt' -exec cp {} /usr/local/share/ca-certificates/ \; && \
      update-ca-certificates; \
    fi && \
    rm -rf /tmp/certificates && \
    apk add --no-cache libpq \
    && addgroup -S app \
    && adduser -S -G app app
```
