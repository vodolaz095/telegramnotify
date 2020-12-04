export app=telegramnotify
export semver=1.0.0
export arch=$(shell uname)-$(shell uname -m)
export gittip=$(shell git log --format='%h' -n 1)
export ver=$(semver).$(gittip).$(arch)
export subver=$(shell hostname)_on_$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
export archiv=build/$(app)-$(arch)-v$(semver)

all: build_dev

deps:
	# install all dependencies required for running application
	go version
	go env

	# installing golint code quality tools and checking, if it can be started
	cd ~ && go get -u golang.org/x/lint/golint
	golint

	# installing golang dependencies using golang modules
	go mod tidy # ensure go.mod is sane
	go mod verify # ensure dependencies are present

lint:
	gofmt  -w=true -s=true -l=true ./
	golint ./...
	go vet ./...

check: lint
	go test -v ./...

build_dev: clean deps check
	go build -o "build/$(app)" main.go

build_prod: clean deps check
#http://www.atatus.com/blog/golang-auto-build-versioning/
	go build -ldflags "-X main.Release="release" -X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app)" main.go
# make binary smaller
	upx build/$(app)

build_without_test:
#http://www.atatus.com/blog/golang-auto-build-versioning/
	go build -ldflags "-X main.Release="release" -X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app)" main.go
# make binary smaller
	upx build/$(app)

build_windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.Release="release" -X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app).exe" main.go
	upx build/$(app).exe

clean:
	go clean

test: check

start:
	go run main.go
