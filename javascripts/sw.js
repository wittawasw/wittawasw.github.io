const CACHE_VERSION = '20240130';

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_VERSION).then((cache) => {
      return cache.addAll([
        '/',
        '/javascripts/sw.js',
        '/assets/css/styles.css',
      ]);
    })
  );
});

self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys().then((cacheNames) => {
      return Promise.all(
        cacheNames.map((cacheName) => {
          if (cacheName !== CACHE_VERSION) {
            return caches.delete(cacheName);
          }
        })
      );
    })
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      // Try to fetch from the network if not in cache
      return (
        response ||
        fetch(event.request).then((fetchResponse) => {
          // Cache the fetched response for future offline access
          return caches.open(CACHE_VERSION).then((cache) => {
            cache.put(event.request, fetchResponse.clone());
            return fetchResponse;
          });
        })
      );
    })
  );
});

// Download and cache the Turbo module
self.addEventListener('fetch', (event) => {
  if (event.request.url.startsWith('https://cdn.skypack.dev/@hotwired/turbo')) {
    event.respondWith(
      fetch(event.request).then((response) => {
        // Cache the Turbo module for offline use
        return caches.open(CACHE_VERSION).then((cache) => {
          cache.put(event.request, response.clone());
          return response;
        });
      })
    );
  }
});
