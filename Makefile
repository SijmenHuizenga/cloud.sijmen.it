
all: build

clean:
	rm -rf -- out

#build a executable
build: deps clean
	cd util && rice embed-go && cd ..
	go build -i -o out/sijmeninstaller

#install all dependencies
deps:
	go get ./...

#build a deployable file
package: build
	upx --brute out/sijmeninstaller

