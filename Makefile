all: build

build:
	mkdir -p bin
	go build -o goflv  goflv.go

clean:
	rm -rf bin
