---
layout: post
title: FizzBuzz First Time.
date: 2016-07-03T17:36:27+09:00
comments: true
categories: ruby programming
cc: FizzBuzz First Time.
keywords: "ruby, programming, fizzbuzz, unit test"
description: "This is my  first time working on FizzBuzz. I've heard of it from
other for some time, didn't pay much attention to it. But an acquaintance asked
me recently how would I do it if I were to be interviewed. So, this's it."
---

> source: [github](https://github.com/wittawasw/fizzbuzz)

This is my  first time working on FizzBuzz. I've heard of it from other for
some time, didn't pay much attention to it. But an acquaintance asked me
recently how would I do if I were to be interviewed and asked to do it.
So, this's it.

Just like when working on any other interview processes, the code should have
quality to ship immediately if needed. I also write along some unit tests in
case I might want to try refactor it after this. Though I'm already feel that
this is already good enough for me.

{% codeblock lang:ruby fizzbuzz.rb%}
module FizzBuzz
  class << self
    def out(range: 1..100)
      range.map { |num| fizzbuzz(num) }
    end

    private

    def fizzbuzz(i)
      if fizzbuzz?(i)
        'fizzbuzz'
      elsif fizz?(i)
        'fizz'
      elsif buzz?(i)
        'buzz'
      else
        i
      end
    end

    def fizzbuzz?(i)
      fizz?(i) && buzz?(i)
    end

    def fizz?(i)
      i % 3 == 0
    end

    def buzz?(i)
      i % 5 == 0
    end
  end
end

# puts FizzBuzz.out.to_s # printing results
{% endcodeblock %}
