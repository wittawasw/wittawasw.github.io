---
layout: post
title: "Backend once and Client everywhere with Rails-API"
date: 2014-04-10 11:41:44 +0700
comments: true
categories: backend rails api angularjs
cc: Backend once and Client everywhere with Rails-API
keywords: "backend, rails, api, angularjs, client, frontend"
description: "Normally when you build a web or mobile application, you will have to build seperate backend because most web framework originally means to generate data along with dynamic HTML templates. But if we think that web application is just a form of native applications and we treat it that way, generate only data and just use static HTML to handle the logic. This way, we can reuse our backend to almost every native applications we want...."
---

Normally when you build a web or mobile application, you will have to build seperate backend because most web framework originally means to generate data along with dynamic HTML templates. But if we think that web application is just a form of native applications and we treat it that way, generate only data and just use static HTML to handle the logic. This way, we can reuse our backend to almost every native applications we want.

>If you're asking for some BIG example out there that use this trick, Try [Twitter.com](https://twitter.com)

Why use Rails ?
---------------
- Huge resources and communities that provide almost everything you want to build applications.
- Very clean Routing and easy to use.
- Secure enough.
- Easiest (in my opinion) to debug using [better_errors](https://github.com/charliesome/better_errors) and [binding_of_caller](https://github.com/banister/binding_of_caller).
- I want to use Rails, that's all.


>Beyond this point require Ruby and Rails installed already.

Getting started: Rails-API
--------------
Assuming that you've already installed Ruby(1.9.3+) and Rails(3.2+) already, today we use this nicely done gem to help us prepare new application

{% highlight shell %}
$ gem install rails-api
$ rails-api new contact-api
{% endhighlight %}

with Rails-API we'll get only essential part of Rails to build API instead of getting all heavy and unnecessary features from traditional Rails. And now we have our newly created app called [contact-api]()

Then put this on your gemfile and run "bundle install", this will make your life a lot easier.

{% highlight ruby%}
group :development do
  gem 'better_errors'
  gem 'binding_of_caller'
end
{% endhighlight %}



Rails time
----------
Since there's only API, so we don't need to setup things like HTML template or homepage here. Let's generate contact by using

{% highlight shell %}
$ rails g scaffold contact name:string email:string
$ rake db:migrate
{% endhighlight %}

If you're running on Rails 3+, you need to put these line in your app/model/contact.rb

{% highlight ruby %}
attr_accessible :name, :email
{% endhighlight %}

If Rails 4+, in app/controllers/contacts_controller

{% highlight ruby %}
def create
  @contact = Contact.new(contact_params)

  if @contact.save
    render json: @contact, status: :created, location: @contact
  else
    render json: @contact.errors, status: :unprocessable_entity
  end
end
private

def contact_params
  params.require(:contact).permit(:name, :email)
end
{% endhighlight %}

and then run your rails server with

{% highlight shell %}
$ rails s
{% endhighlight %}

Try access by your browser to [localhost:3000/contacts](http://localhost:3000/contacts/) and you will see an empty JSON "[]" there. Now let's try to make it error by access to [localhost:3000/contacts/test](http://localhost:3000/contacts/test) and you'll see error like below.
![contact error](/images/post/contact-error.png)

Obviously, it's a "Not found" error because we try to access object Contact that have id = "test" which doesn't exist. On the right side, you'll see a very handy shell that you can play with it like how you play in "rails console". And since we have our console right before our eyes, we should try something.

{% highlight ruby %}
Contact.create!(name:"John", email:"fathername@example.com")
{% endhighlight %}

and result is

![contact create](/images/post/contact-create.png)

Now we can try to access [localhost:3000/contacts/1](http://localhost:3000/contacts/1) and you will see JSON of your newly created contact.
![contact 1](/images/post/contact-1.png)


Prepare for CORS
----------------
What we trying to do here is to actually access to Rails server([localhost:3000](http://localhost:3000)) from another server instance. So it'll inevitably face with [CORS](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing) problem and here we'll cover that problem before it's occured with [rack-cors](https://github.com/cyu/rack-cors)

Put this line in Gemfile

{% highlight ruby%}
    gem 'rack-cors', :require => 'rack/cors'
{% endhighlight %}

then

{% highlight ruby %}
    bundle install
{% endhighlight %}

Finally, put this code in config/application.rb inside class Application and it'll look like this.

{% highlight ruby %}
module ContactApi
  class Application < Rails::Application
    #.....
    config.middleware.use Rack::Cors do
      allow do
        origins '*'
        resource '*', :headers => :any, :methods => [:get, :post, :options]
      end
    end
  end
end
{% endhighlight %}

Restart Rails server and it'll become accessible from every origin.




HTML time
---------
It's time to build our static site to serve as a native client for contact-api. But before that, we have to choose which server to serve our static files. It can be many things such as another Rails, Node.js, PHP, etc. Whichever is fine but I choose to go with [Go](http://golang.org), no reason, I just want it. (See details of how easy Go handle static file in [previous post](http://wittawasw.com/blog/2014/04/09/free-and-easy-static-website-hosting-using-google-app-engine-and-golang/))

Below is my ~~jsFiddle~~(obsoleted) demo using angularjs as a client to add contact object by using Rails backend as API only. Or you can play a demo at this Google App Engine [Site](http://rails-angular-backend-demo.appspot.com/)

All the code can be found at [Github](https://github.com/wittawasw/rails-angular-backend-demo)


