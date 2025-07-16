---
layout: post
title: บังคับให้ browser tab อยู่ในสภาพ active
tags: javascript, browser
date: 2025-07-16 10:16 +0700
---
### แก้ที่ API พื้นฐาน browser console

```javascript
// พิมพ์ลงไปได้เลย  browser console
Object.defineProperty(document, 'hidden', { get: () => false });
Object.defineProperty(document, 'visibilityState', { get: () => 'visible' });

window.addEventListener('visibilitychange', (e) => {
  e.stopImmediatePropagation();
}, true);

```

### ทำเป็น Bookmark link ไว้กด

```javascript
// wrap เป็น function ก่อน กับใส่ alert
javascript:(function() {
  Object.defineProperty(document, 'hidden', { get: () => false });
  Object.defineProperty(document, 'visibilityState', { get: () => 'visible' });
  window.addEventListener('visibilitychange', e => e.stopImmediatePropagation(), true);
  alert("Tab visibility spoof activated");
})();

// encode, ตรงนี้ทำใน browser console หรือจะ run ด้วย nodejs ก็ได้
const s = `
javascript:(function() {
  Object.defineProperty(document, 'hidden', { get: () => false });
  Object.defineProperty(document, 'visibilityState', { get: () => 'visible' });
  window.addEventListener('visibilitychange', e => e.stopImmediatePropagation(), true);
  alert("Tab visibility spoof activated");
})();
`
// เอา link ไปใส่ใน URL ตอนสร้าง bookmark
const link = encodeURIComponent(s)
```

### Bookmarkable script

```javascript
javascript:(function()%7BObject.defineProperty(document%2C%20'hidden'%2C%20%7B%20get%3A%20()%20%3D%3E%20false%20%7D)%3BObject.defineProperty(document%2C%20'visibilityState'%2C%20%7B%20get%3A%20()%20%3D%3E%20'visible'%20%7D)%3Bwindow.addEventListener('visibilitychange'%2C%20function(e)%20%7Be.stopImmediatePropagation()%3B%7D%2C%20true)%3Balert(%22Tab%20visibility%20spoof%20activated%22)%3B%7D)()
```
