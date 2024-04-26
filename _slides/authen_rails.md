---
marp: true
theme: default
class: invert
---

<style>
  @import url('https://fonts.googleapis.com/css2?family=Anuphan:wght@100..700&display=swap');

  :root {
    --color-1: #D30001;
    --color-2: #F0F0F4;
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
    color: var(--color-1) !important;
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
</style>

# **Backend Authentication ด้วย Ruby on Rails**
## 26 เมษายน 2567

---
# ทำความเข้าใจกับ Backend Authentication ?

---
# Ruby on Rails

- MVC
- Convention over Configuration


---
# HTTP Basic Access Authenentication


---
# Current

[https://api.rubyonrails.org/classes/ActiveSupport/CurrentAttributes.html](https://api.rubyonrails.org/classes/ActiveSupport/CurrentAttributes.html)

```ruby
# app/models/current.rb
class Current < ActiveSupport::CurrentAttributes
  attribute :account, :user
  attribute :request_id, :user_agent, :ip_address

  resets { Time.zone = nil }

  def user=(user)
    super
    self.account = user.account
    Time.zone    = user.time_zone
  end
end
```

---
# Username + Secured Password

---
# Bcrypt

```ruby
require 'bcrypt'

class User < ActiveRecord::Base
  # users.password_hash in the database is a :string
  include BCrypt

  def password
    @password ||= Password.new(password_hash)
  end

  def password=(new_password)
    @password = Password.create(new_password)
    self.password_hash = @password
  end
end
```

---
# Cookie and Session

It's called a "cookie" because it's a term borrowed from HTTP cookies, which are small pieces of data stored on the client's browser. In Rails, the term "cookie" refers to the data stored in the client's browser, similar to HTTP cookies.

a magic cookie, or just cookie for short, is a token or short packet of data passed between communicating programs.

The concept of sessions in web development was introduced around the mid-1990s. The exact date of invention might vary depending on the specific technology or framework being used.

---
# Cookie sv JWT Token

    JWT (JSON Web Token):
        Token-based authentication mechanism.
        Data is stored in a token, which is sent to the client.
        Stateless: Server doesn't need to store session data.
        Typically used in stateless APIs.

    Cookie:
        Data is stored on the client-side.
        Sent with every HTTP request to the server.
        Can be used for session management in web applications.
        Stateful: Server may need to store session data.

        JWT can be secure if used correctly. However, it's important to consider the type of data being stored and the security measures in place. For sensitive data, using server-side session management (e.g., cookies with server-side session storage) might be preferable because it gives more control over the data and reduces the risk of tampering.

---
# JWT

JWT signatures are created using cryptographic algorithms, such as HMAC (Hash-based Message Authentication Code) or RSA (Rivest-Shamir-Adleman).

For HMAC:

    A secret key is used to generate the signature.
    The signature is a hash (e.g., SHA256) of the header and payload, encoded using the secret key.

For RSA:

    A public/private key pair is used.
    The signature is created using the private key and verified using the public key.

The signature ensures the integrity of the JWT and prevents tampering.

---
# JWT

https://github.com/jwt/ruby-jwt


---
# Session Token vs JWT Token

---
# Signed Token

---
# One Time Password

---
# Log Filtering
- https://guides.rubyonrails.org/action_controller_overview.html#log-filtering


---
# OAuth 2.0

---
# Fundamentals of Authentication and Authorization in Rails

## Description

Explain the core concepts of authentication and authorization in Ruby on Rails, gaining an understanding of their fundamental principles and implementation techniques. This talk will guide you through the process of building secure and efficient authentication and authorization systems in Rails applications.

Explore key topics such as secure session management, password hashing, role-based access control, and integration with web/mobile clients. This will feature a code-walkthrough
