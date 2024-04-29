web: bundle exec jekyll s

# watch_slides: while inotifywait -q -e close_write _slides; do make build_slides; done
watch_slides: while [ "$(uname)" = "Darwin" ]; do fswatch -1 _slides && make build_slides; done || while inotifywait -q -e close_write _slides; do make build_slides; done
