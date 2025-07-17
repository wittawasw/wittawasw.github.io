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

### Edited: เพิ่มการขยับ mouse, click

```javascript
javascript:(function() {
  Object.defineProperty(document, 'hidden', { get: () => false });
  Object.defineProperty(document, 'visibilityState', { get: () => 'visible' });
  window.addEventListener('visibilitychange', e => e.stopImmediatePropagation(), true);

  function simulateClickAndMove() {
    const el = document.querySelector('.header');
    if (el) el.click();

    const evt = new MouseEvent('mousemove', {
      bubbles: true,
      clientX: Math.random() * window.innerWidth,
      clientY: Math.random() * window.innerHeight
    });
    document.dispatchEvent(evt);
  }

  setInterval(simulateClickAndMove, 30000); // 30 วิ
  alert("Tab visibility spoof + periodic click/move started");
})();
```

```javascript
// Bookmark
javascript:(function()%7BObject.defineProperty(document%2C%20'hidden'%2C%20%7B%20get%3A%20()%20%3D%3E%20false%20%7D)%3BObject.defineProperty(document%2C%20'visibilityState'%2C%20%7B%20get%3A%20()%20%3D%3E%20'visible'%20%7D)%3Bwindow.addEventListener('visibilitychange'%2C%20e%20%3D%3E%20e.stopImmediatePropagation()%2C%20true)%3Bfunction%20simulateClickAndMove()%20%7Bconst%20el%20%3D%20document.querySelector('.classroom-workspace-overview__header')%3Bif%20(el)%20el.click()%3Bconst%20evt%20%3D%20new%20MouseEvent('mousemove'%2C%20%7Bbubbles%3A%20true%2C%20clientX%3A%20Math.random()*window.innerWidth%2C%20clientY%3A%20Math.random()*window.innerHeight%7D)%3Bdocument.dispatchEvent(evt)%3B%7DsetInterval(simulateClickAndMove%2C%2030000)%3Balert('Tab%20visibility%20spoof%20%2B%20periodic%20click%2Fmove%20started')%3B%7D)()
```
