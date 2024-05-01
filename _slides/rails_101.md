---
marp: true
theme: default
class: invert
---

<script type="module">
  import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
  mermaid.initialize({ startOnLoad: true, theme: 'dark' });
</script>

<style>
  @import url('https://fonts.googleapis.com/css2?family=Anuphan:wght@100..700&display=swap');

  :root {
    --color-1: #D30001;
    --color-2: #CE4D4D;
    --color-3: #F0F0F4;
  }

  section {
    background: #1E1C1C;
  }

  h1, b, strong,
  h2, h3, h4,
  a,
  p {
    font-family: "Anuphan", sans-serif;
  }

  h1, b, strong {
    color: var(--color-1) !important;
    font-weight: bold;
  }

  h2, h3, h4 {
    color: var(--color-2) !important;
    font-weight: 500;
  }

  a {
    color: var(--color-2) !important;
    font-weight: 500;
    font-style: italic;
    text-decoration: underline;
  }

  p {
    color: var(--color-3) !important;
    font-weight: 300;
  }

  footer {
    font-size: 20px;
    text-align: right;
  }

  .mermaid {
    position: relative;
    left: 15vw;
  }

  .large {
    position: relative;
    left: 0vw !important;
  }
</style>

# **Ruby on Rails: 101**
## 1 พฤษภาคม 2567

---
# ติดตั้ง Ruby
<!--
เช็คว่าผู้เข้าร่วมติดตั้งของจำเป็นแล้วก็สามารถใช้งานได้
-->
  ```sh
  # Ruby Version Manager
  brew install rbenv

  # Latest list of Ruby versions
  brew install ruby-build

  # Use ENV in Ruby
  brew install rbenv-vars
  ```
  ```sh
  which ruby
  #=> /.rbenv/shims/ruby

  # rbenv-vars
  rbenv vars
  ```

---
# ติดตั้ง Rails
<!--
เช็คว่าผู้เข้าร่วมติดตั้งของจำเป็นแล้วก็สามารถใช้งานได้
-->
  ```sh
  gem install bundler
  gem install rails
  ```

---
# ทบทวน

- Web Server
- HTTP
- HTML, JS, CSS
- API, json
- Testing

---

# Overview

- คำสั่ง `rails`
- โครงสร้างไฟล์
-

---

# คำสั่ง `rails`

- bundle exec
- new
- server
- generate
- test

---
# โครงสร้างของไฟล์

---
# Ruby on Rails

<!--
  How Rails stands out of other web frameworks.
  A progenitor of modern full-stack web.

  Talk about the approach that we will see when showing
  files from Rails.
-->

- MVC Web Framework
- Convention over Configuration

```javascript
// models
user.rb

// views
users/index.html
users/show.html

// controllers
users_controller.rb

// services
user_service.rb
```

---
