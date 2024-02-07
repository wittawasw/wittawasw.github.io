const CACHE_VERSION = '202401300614';

const urlsToCache = [
  '/',
  
  '/404.html',
  
  '/about.html',
  
  '/slides/code_reviews.html',
  
  '/slides/examples.html',
  
  '/javascripts/sw.js',
  
  '/assets/css/style.css',
  
  '/assets/minima-social-icons.svg',
  
  '/index.html',
  
  '/page2/index.html',
  
  '/tag/startup/index.html',
  
  '/tag/gpt/index.html',
  
  '/tag/rails/index.html',
  
  '/tag/turbo/index.html',
  
  '/tag/hotwired/index.html',
  
  '/tag/html/index.html',
  
  '/tag/javascript/index.html',
  
  '/tag/pdf/index.html',
  
  '/tag/converter/index.html',
  
  '/tag/imagemagick/index.html',
  
  '/collection/posts/index.html',
  
  '/collection/posts/page2/index.html',
  
  '/feed.xml',
  
  '/sitemap.xml',
  
  '/robots.txt',
  
  
  '/2024/01/22/convert-pdf-grayscale-ubuntu.html',
  
  '/2024/01/08/%E0%B8%A5%E0%B8%AD%E0%B8%87%E0%B9%83%E0%B8%8A%E0%B9%89-hotwired-turbo-%E0%B9%83%E0%B8%99%E0%B9%80%E0%B8%A7%E0%B9%87%E0%B8%9A%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%83%E0%B8%8A%E0%B9%88-rails.html',
  
  '/2023/02/22/%E0%B8%95%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B8%E0%B8%99%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9E%E0%B8%B1%E0%B8%92%E0%B8%99%E0%B8%B2%E0%B8%9C%E0%B8%A5%E0%B8%B4%E0%B8%95%E0%B8%A0%E0%B8%B1%E0%B8%93%E0%B8%91%E0%B9%8C%E0%B9%81%E0%B8%9A%E0%B8%9A-ChatGPT.html',
  
  '/2022/04/25/use-docker-in-local-web-application-development.html',
  
  '/2022/02/20/comment-in-code.html',
  
  '/2022/02/08/resume-%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%84%E0%B8%94%E0%B9%89%E0%B8%87%E0%B8%B2%E0%B8%99.html',
  
  '/2022/01/31/%E0%B9%81%E0%B8%99%E0%B8%A7%E0%B8%84%E0%B8%B4%E0%B8%94%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AA%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%AA%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B8%A1.html',
  
  '/2022/01/27/%E0%B8%99%E0%B8%B4%E0%B8%A2%E0%B8%B2%E0%B8%A1%E0%B8%82%E0%B8%AD%E0%B8%87-startup.html',
  
];

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_VERSION).then((cache) => {
      return cache.addAll(urlsToCache);
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
        fetch(event.request).then((fetchResponse) => {
          // Cache the fetched response for future offline access
          caches.open(CACHE_VERSION).then((cache) => {
            cache.put(event.request, fetchResponse.clone());
          });
          return fetchResponse;
        }).catch(() => {
          // If network fetch fails, use the cached response
          return response || caches.match('/');
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
        caches.open(CACHE_VERSION).then((cache) => {
          cache.put(event.request, response.clone());
        });
        return response;
      })
    );
  }
});
