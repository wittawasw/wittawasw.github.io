---
layout: post
title: Easy HTTPS Proxy Using Caddy on DigitalOcean.
date: 2017-04-28T02:11:27+07:00
comments: true
categories: https proxy digitalocean caddy webserver docker
cc: Easy HTTPS Proxy Using Caddy on DigitalOcean.
keywords: "webserver, proxy, digitalocean, caddy, https, docker"
description: "If you need a good HTTPS proxies that come with auto Let's Encrypt
certificates and already a user of DigitalOcean. Here I've got an easy but solid
way to deploy one. Already using this on a production servers with no problem
so far."
---

> The setup need DigitalOcean account and Domain on DigitalOcean's nameserver.
Otherwise, it'd need a different step to complete.

If you need a good HTTPS proxies that come with auto [Let's Encrypt](https://letsencrypt.org/)
certificates and already a user of DigitalOcean. Here I've got an easy but solid
way to deploy one. Already using this on a production servers with no problem
so far.

> [Caddy](https://caddyserver.com/) server is already the easiest server that
you can [install](https://caddyserver.com/tutorial) on Linux. But I decided to
use docker with it because it's easier for me to control volume and settings.

### From beginning to deploy.
- Setup domain to use DigitalOcean's nameserver.
- Create droplet using pre-defined `docker` droplet.
- Add `A` or `CNAME` that point domain to droplet's IP address.
- `SSH` into droplet and create directory `caddy` in home directory.
- Enable `ufw` port `443`.
- Create DigitalOcean auth token and keep it close.
- Create `Caddyfile` in home directory.
- Create `docker-compose.yml` in home directory.
- From home directory, use command `docker-compose up -d`
- Caddy's online. That's it. No more hassle.

{% highlight YAML %}
  # proxy's url (domain that we set record earlier.)
front-end-url:443 {
  # backend's servers can be single or multiple servers
  # SSL terminate here.
  proxy / backend-url-1:80 backend-url-2:80 backend-url-3:80 {
    policy round_robin
    header_upstream Host {host}
    header_upstream X-Real-IP {remote}
    header_upstream X-Forwarded-For {remote}
    header_upstream X-Forwarded-Proto {scheme}
  }

  gzip

  tls {
   # verify ssl certificate using dns from DigitalOcean really save me from a lot of pain.
   dns digitalocean
  }
}

{% endhighlight %}

{% highlight YAML %}
version: '3'
services:
  proxy:
    image: zzrot/alpine-caddy
    ports:
      -   443:443
    volumes:
      -   ./Caddyfile:/etc/Caddyfile
      -   ./caddy:/root/.caddy
      # don't forget to add volume, otherwise you will request new certificate
      # every time server restart and exceed letsencrypt quota in the process.
    environment:
      - DO_AUTH_TOKEN=<DO_ACCOUNT_TOKEN>
      # get this token from setting page.
{% endhighlight %}

### Where to go from here
See official [documentation](https://caddyserver.com/docs) of Caddy. There's a
lot of features that can be easily setup using simple directives.
