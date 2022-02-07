FROM ruby:3.0.3-alpine as jekyll

RUN apk add --no-cache build-base gcc bash cmake git

RUN gem install bundler -v "~>2.3.6"
COPY Gemfile Gemfile.lock /
RUN bundle install --jobs 10

# /app is mounted in docker-compose
WORKDIR /app

EXPOSE 4000

CMD [ "bundle", "exec", "jekyll", "serve", "--force_polling"]
