default: build

build:
	go build -o bin/git-pull-requests

clean:
	rm -rf bin/git-pull-requests
