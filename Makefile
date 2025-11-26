build:
	go build

dev:
	fd --type file -e go -e templ --exclude '*_templ.go' | entr -r sh -c "tailwindcss -i ./static/css/input.css -o ./static/css/output.css && templ generate && go run ."

