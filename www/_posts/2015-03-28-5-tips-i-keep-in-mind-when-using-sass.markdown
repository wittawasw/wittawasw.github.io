---
layout: post
title: "5 tips I keep in mind when using Sass"
date: 2015-03-28T02:10:53+07:00
comments: true
categories: Sass stylesheet tips
keywords: "Sass, stylesheet, tips"
description: "It's been some time already since I've been introduced to Sass, Though I normally use just Scss syntax. Sass's really nice additional tool and these tips are what I found myself comfortable with and tend to convince people around me to use..."
header-img: "images/post/04-sass.jpg"
---

>If you still don't know what exactly is Sass, here is [its official site](http://Sass-lang.org)

It's been some time already since I've been introduced to Sass, Though I normally use just Scss syntax. Sass's really nice additional tool and these tips are what I found myself comfortable with and tend to convince people around me to use.


*******


1.Structure is important
-----------------------
Since I'm mostly using Rails, sometimes people tend to use Rails's default structure of stylesheet files base on controllers which is not wrong but I think it's less useful than it should be.

{% codeblock %}
- app
  + assets
    * stylesheets
      - application.scss
      - post.scss
      - comment.scss
      - user.scss
{% endcodeblock %}

But it's not reusable and just think about when this project grow up to 20+ controllers and 50+ pages is enough to makes me feels dizzy.


My approach is to divide Sass files into modules and widgets based on how it's used.

{% codeblock %}
- app
  + assets
    * stylesheets
      - application.scss
      - variables.scss
      - layouts
        * navigation_bar.scss
        * footer.scss
        * container.scss
      - partials // use '_' in front of partials
        * _all.scss
        * _button.scss
        * _grid.scss
        * _bar.scss
        * _typography.scss
      - modules // divide functionality code into modules
        * forms
          - register_form.scss
          - post_form.scss
          - comment_form.scss
        * cart
          - ...
{% endcodeblock %}

&nbsp;


2.Keep variables together
------------------------

A nice feature of Sass that every beginner would know about is how Sass cleverly handle values in form of variables. But to put it all over every files is not practical since I will have troubles finding them in later stage of development. So, this approach is simple, just try to lump it together as much as I can. In my most cases, one file is enought.

{% codeblock lang:scss variables.scss%}
// Our variables
$base-font-family: Helvetica, Arial, sans-serif;
$base-font-size:   16px;
$small-font-size:  $base-font-size * 0.875;
$base-line-height: 1.5;

$spacing-unit:     30px;

$text-color:       #111;
$background-color: #fdfdfd;
$brand-color:      #2a7ae2;

$grey-color:       #828282;
$grey-color-light: lighten($grey-color, 40%);
$grey-color-dark:  darken($grey-color, 25%);
{% endcodeblock %}

******
&nbsp;


3.Order import statements
-----------------------

Another features of Sass that's widely known, **@import**.
This may sound insignificant at first but to not carefully ordering import statements can prove to be troubled later and it will also provide more readablity to our stylesheets. So, it's a win in every way.

{% codeblock lang:scss application.scss%}
@import 'variables';

@import 'partials/button';
@import 'partials/grid';
@import 'partials/bar';
@import 'partials/typography';

@import 'layouts/navigation_bar';
@import 'layouts/footer';
@import 'layouts/container';

@import 'modules/forms';
@import 'modules/cart';

{% endcodeblock %}
******
&nbsp;

4.Loop and Conditions are handy.
---------------------------------
> A nice Sass playground [here](http://Sassmeister.com)


It's similar to every programming languages you've used so far but it actually in stylesheet files. Learn to use it and you'll never want to go back to plain css again.

Here's an example.

{% codeblock lang:scss _grid.scss%}
// SCSS
@for $i from 1 through 12 {
  %columns-#{$i} {
    $width: ((100 / 12) * $i) * 1%;
    @if $width == 100% {
      $width: 99%;
    }
    width: $width;
  }
}

@for $i from 1 through 12 {
  .span-#{$i} {
    @extend %columns-#{$i};
  }
}


// Compiled CSS

.span-1 {
  width: 8.33333%;
}
.span-2 {
  width: 16.66667%;
}
.span-3 {
  width: 25%;
}
.span-4 {
  width: 33.33333%;
}
.span-5 {
  width: 41.66667%;
}
.span-6 {
  width: 50%;
}
.span-7 {
  width: 58.33333%;
}
.span-8 {
  width: 66.66667%;
}
.span-9 {
  width: 75%;
}
.span-10 {
  width: 83.33333%;
}
.span-11 {
  width: 91.66667%;
}
.span-12 {
  width: 99%;
}


{% endcodeblock %}

&nbsp;
    

5.Use @include, @extend wisely
-----------------------------

I found this a lot in past year. When people know how to use @mixin, they failed to notice how large their compiled CSS has become. Sometimes it's not really an issue since I can have almost or perhaps the same results when I use only @include with @mixin. But my compiled CSS file is not nearly the same though.

Let try something to see it in action,


### Use @include to include @mixin properties
{% codeblock lang:scss include.scss%}

// SCSS
@mixin icon($social) {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-size: 77px 35px;
  background-repeat: no-repeat;
  background-image:url($social);
}
 
.icon-facebook {
  @include icon('/img/icon-facebook.png');
}
 
.icon-twitter {
  @include icon('/img/icon-twitter.png');
}

.icon-googleplus {
  @include icon('/img/icon-googleplus.png');
}

// Compiled CSS
 
.icon-facebook {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-size: 77px 35px;
  background-repeat: no-repeat;
  background-image:url('/img/icon-facebook.png');
}
 
.icon-twitter {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-size: 77px 35px;
  background-repeat: no-repeat;
  background-image:url('/img/icon-twitter.png');
}

.icon-googleplus {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-size: 77px 35px;
  background-repeat: no-repeat;
  background-image:url('/img/icon-googleplus.png');
}
{% endcodeblock %}
&nbsp;

### Use @extend to extend properties
{% codeblock lang:scss extend.scss%}

// SCSS
%icon {
  display: inline-block;
  width: 14px;
  height: 14px;
  background-size: 77px 35px;
  background-repeat: no-repeat;
}
 
.icon-facebook {
  @extend %icon;
  background-image:url('/img/icon-facebook.png');
}
 
.icon-twitter {
  @extend %icon;
  background-image:url('/img/icon-twitter.png');
}

.icon-googleplus {
  @extend %icon;
  background-image:url('/img/icon-googleplus.png');
}

// Compiled CSS
.icon-facebook, .icon-twitter, .icon-googleplus {
  display: inline-block
  width: 14px
  height: 14px
  background-size: 77px 35px
  background-repeat: no-repeat;
}
 
.icon-facebook {
  background-image:url('/img/icon-facebook.png');
}
 
.icon-twitter {
  background-image:url('/img/icon-twitter.png');
}

.icon-googleplus {
  background-image:url('/img/icon-googleplus.png');
}
{% endcodeblock %}


Both compiled CSS give us 26 and 19 lines with the same results on browser. There's still differrences in size though it's not much to the point that it doesn't matter but it's still our choices to choose which approach to do it.In this case, I choose latter choice to save some spaces for compiled css, more reasonable to me.
&nbsp;

