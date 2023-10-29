build:
	go build -ldflags="-s -w"

install:
	go install
	
run:
	go build && ./svault

mod_init:
	go mod init github.com/ragul28/svault
	go get -u

mod:
	go mod tidy
	go mod verify
	go mod vendor
