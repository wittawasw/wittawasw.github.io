---
layout: post
title: keep-browser-tab-active
---

```javascript
Object.defineProperty(document, 'hidden', { get: () => false });
Object.defineProperty(document, 'visibilityState', { get: () => 'visible' });

window.addEventListener('visibilitychange', (e) => {
  e.stopImmediatePropagation();
}, true);
```
