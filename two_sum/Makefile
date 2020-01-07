test:
	$(info [Running Unit Tests])
	cd solution && go test

build:
	rm -rf Gopkg.lock
	dep ensure
	cd solution && go test && cd ..
	env GOOS=linux go build -ldflags="-s -w" -o bin/solution solution/main.go