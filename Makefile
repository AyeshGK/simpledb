build:
	go build -o bin/simpledb src/main.go

run:
	./bin/simpledb

test:
	go test -v ./...

clean:
	rm -rf bin

# make clean build run 
cbr: clean build run