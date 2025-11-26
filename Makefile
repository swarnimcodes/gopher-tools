templ:
	templ generate

css:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css 

gobuild:
	go build

build: css templ gobuild

run: build
	./gopher-toolbox

dev:
	fd --type file -e go -e templ --exclude '*_templ.go' | entr -r sh -c "tailwindcss -i ./static/css/input.css -o ./static/css/output.css && templ generate && go run ."


