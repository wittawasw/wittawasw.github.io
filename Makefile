deploy:
	bundle exec jgd -r main

build_about:
	npx @marp-team/marp-cli aboutme.md
	npx @marp-team/marp-cli aboutme.md --pdf

build_slides:
	npx @marp-team/marp-cli -I _slides -o slides --html
