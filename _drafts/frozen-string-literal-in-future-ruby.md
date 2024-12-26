---
layout: post
title: frozen string literal in future ruby
---

ruby 3.4.1 (2024-12-25 revision 48d4efcb85) +PRISM [x86_64-linux]
➜  ~ ruby -W:deprecated test.rb
test.rb:2: warning: literal string will be frozen in the future (run with --debug-frozen-string-literal for more information)
asdasdasd
➜  ~ ruby -W:deprecated test.rb --debug-frozen-string-literal
test.rb:2: warning: literal string will be frozen in the future (run with --debug-frozen-string-literal for more information)
asdasdasd
➜  ~ ruby -W:deprecated --debug-frozen-string-literal test.rb
test.rb:2: warning: literal string will be frozen in the future
test.rb:1: info: the string was created here
asdasdasd
