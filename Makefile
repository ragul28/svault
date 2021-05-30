build:
	GOOS=linux GOARCH=amd64 go build

install:
	go install
	
run:
	go build && ./svault

init:
	go mod init github.com/ragul28/svault
	go get -u

mod:
	go mod tidy
	go mod verify
	go mod vendor
