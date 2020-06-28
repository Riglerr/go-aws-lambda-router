.PHONY: build publish
build:
	GOARCH=amd64 GOOS=linux go build ./... 

build-debug:
	GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o lambda
	mkdir -p .aws-sam/build/RouterFunction/
	cp lambda .aws-sam/build/RouterFunction/
run-debug:
	sam local invoke -d 5986 --debugger-path /home/rob/go/bin --debug-args "-delveAPI=2" -e e.json

run: 
	sam local invoke -e e.json
publish:
	aws lambda update-function-code --function-name router --zip-file fileb://main.zip