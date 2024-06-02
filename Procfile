web: bundle exec jekyll s
watch_slides: if [ "$(uname)" = "Darwin" ]; then fswatch -1 _slides && make build_slides; else while inotifywait -q -e close_write _slides; do make build_slides; done; fi
