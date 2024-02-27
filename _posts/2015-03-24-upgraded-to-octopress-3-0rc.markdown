---
layout: post
title: "Upgraded to Octopress 3.0rc"
date: 2015-03-24T13:53:32+07:00
comments: true
tags: jekyll ruby octopress
cc: Upgraded to Octopress 3.0rc
keywords: "jekyll, ruby, octopress, bootstrap"
description: "It's been almost a year since I last updated this blog, to me, so many things happened and things are going to flow a lot faster this year. I've decided to do this since last month but I want to upgrade this blog from Octopress 2.0 -> 3.0 and I just didn't feel like I want to do it at that time since I thought it'd mean a lot of works to do the upgrade...."
---

>I'm too lazy to do tutorial on this one, but you can look at the source [here](https://github.com/wittawasw/wittawasw.github.io/tree/source) directly.

It's been almost a year since I last updated this blog, to me, so many things happened and things are going to flow a lot faster this year. I've decided to do this since last month but I want to upgrade this blog from Octopress 2.0 -> 3.0 and I just didn't feel like I want to do it at that time since I thought it'd mean a lot of works to do the upgrade. I was wrong.

It's clearly stated in [Octopress's announcement](http://octopress.org/2015/01/15/octopress-3.0-is-coming/) about major upgrade to 3.0 that it'd re-engineer the whole things. By turning Octopress into just Jekyll cli tools, people won't have to learn the very similar things twice anymore when going from Jekyll to Octopress, It's basically just Jekyll with a nice cli tool in Octopress 3.0

I started doing this for a while and found that it's much more easier than when I was struck with Octopress 2.0. Migrating is easy, I just moved my **_posts** folder over to this new repo and it's done.

But some plugins will not work immedietly but it's just a matters of relocating Ruby lib files and debugging some small errors that happened.

And with the new site in place I've decided to change a bit of front-end myself by using [Bootstrap 3.x](http://getbootstrap.com/) and adding some cover to make it more appealing.

Right now, I'm done with this. If anyone still reluctant to upgrade your Octopress, I really recommend to do it. It's worth your time since it won't take long.

