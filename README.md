## Development with Ruby enviroment
- Ruby 2.7+ (lower version may works but they're not being tested)
- `gem install bundler` to install bundler.
- `bundle install` to install development dependencies.
- `bundle exec jekyll s` to start development server.
- `bundle exec jekyll s --draft` to start development server with drafted posts.
- `bundle exec jgd -r [SOURCE_BRANCH]` for deployment to Github pages,
  when `SOURCE_BRANCH` is a branch in git that contains jekyll source code.

## Development with Docker Compose
- `docker-compose up` or `docker-compose run` for daemon process

Using `docker` either standalone or with `docker-compose` will results in
making jekyll cache directory `.jekyll-cache` ended up with docker's permission
and may have `permission denied` error if trying to open again with local ruby
environment. So, it's better to choose one and stick to it without changing.
