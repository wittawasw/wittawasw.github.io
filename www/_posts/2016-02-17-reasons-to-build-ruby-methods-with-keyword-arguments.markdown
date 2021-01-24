---
layout: post
title: Reasons to Build Ruby Methods with Keyword Arguments
date: 2016-02-17T02:08:05+07:00
categories: ruby method argument
cc: Reasons to Build Ruby Methods with Keyword Arguments
keywords: "ruby, method, argument"
description: "Keyword Arguments make building and maintaining Ruby methods
easier and clearer than normal method declaration."

---

> Keyword Arguments make building and maintaining Ruby methods easier and
clearer than normal method declaration.


Keyword Arguments in Ruby is one of the good things that my colleages tend to
ignore by not using it. Despite that it doesn't have any proper downside of
using it and come with many good points.

### It make method clearer at first glance.
This is obvious. I don't have to look for source of method when I see the
method being called at other places.
{% highlight ruby linenos %}
# normal initializer
class Post
  def initialize(subject, body)
    @subject = subject
    @body = body
  end
end

# you will have to find the above Class to understand just this line.
Post.new('hello', 'this is a msg')
{% endhighlight  %}

compare to how Keyword Arguments work

{% highlight ruby linenos %}
# initializer with keywords
class Post
  def initialize(subject:, body:)
    @subject = subject
    @body = body
  end
end

# likely to understand the basic functionality at firstglance
Post.new(subject: 'hello',
         body: 'this is a msg',
        )
{% endhighlight  %}

&nbsp;


### Better understanding error message.
Keyword Arguments methods provide better error message when thing gone wrong.

{% highlight shell %}
# Post.new('hello')
ArgumentError: wrong number of arguments(given 1, expected 2)

# Post.new(title: 'hello')
ArgumentError: missing keyword: body
{% endhighlight %}

&nbsp;


### Default setter and required key
This is indeed useful, for you can quickly look at given params and understand
what needs to be done almost immediately.
{% highlight ruby %}
class Post
  # default setter for :title to be 'untitled' if nil
  # required key :body if leave the setter blank.
  def initialize(title: 'untitled', body:)
    @title = title
    @body = body
  end
end)
{% endhighlight %}

&nbsp;

