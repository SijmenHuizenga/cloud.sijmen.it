
all: cleanup downloader make-install-os make-setup-os

cleanup:
	rm -r out
	mkdir out

downloader:
	cp download.sh out/download.sh

make-install-os:
	cp fdisk.layout out/fdisk.layout
	env GOOS=linux GOARCH=amd64 go build -o out/install-os install-os.go util.go

make-setup-os:
	env GOOS=linux GOARCH=amd64 go build -o out/setup-os setup-os.go util.go
