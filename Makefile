export app=telegramnotify
export majorVersion=1
export minorVersion=0

export patchVersion=$(shell git log origin/master --format='%h' | wc -l)
export semver=$(majorVersion).$(minorVersion).$(shell git log origin/master --format='%h' | wc -l)
export arch=$(shell uname)-$(shell uname -m)
export gittip=$(shell git log --format='%h' -n 1)
export ver=$(semver).$(gittip).$(arch)
export subver=$(shell hostname)_on_$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
#export subver='Code compiled by $(shell hostname) on $(shell date -u '+%Y-%m-%d_%HH:%MM:%SS')'
export archiv=build/$(app)-$(arch)-v$(semver)

all: sign

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
	go build -ldflags "-X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app)" main.go
	upx build/$(app)
	./build/telegramnotify --version

build_without_test:
#http://www.atatus.com/blog/golang-auto-build-versioning/
	go build -ldflags "-X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app)" main.go
	upx build/$(app)

build_windows: clean deps check
# architecture can be changed here to make cross compilation work properly
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app).exe" main.go
	upx build/$(app).exe

build_macos: clean deps check
# architecture can be changed here to make cross compilation work properly
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X github.com/vodolaz095/telegramnotify/commands.Subversion=$(subver) -X github.com/vodolaz095/telegramnotify/commands.Version=$(ver)" -o "build/$(app)_macos" main.go
	upx build/$(app)_macos


build_rpm: build_prod
	rpmdev-wipetree
	rpmdev-setuptree
	cp build/$(app) $${HOME}/rpmbuild/BUILD/$(app)
	cp contrib/telegramnotify.json $${HOME}/rpmbuild/BUILD/$(app)
	rpmbuild --clean -bb --define 'release ${gittip}' --define 'version $(majorVersion).$(minorVersion).$(patchVersion)' $(app).spec
	cp $${HOME}/rpmbuild/RPMS/$(shell uname -m)/$(app)-$(majorVersion).$(minorVersion).$(patchVersion)-$(gittip).$(shell uname -m).rpm build/


sign: build_rpm build_windows build_macos
	rm build/*.txt -f
	rm build/*.txt.sig -f
	find build/ -name "$(app)*" -exec md5sum {} + > build/md5sum.txt
	gpg2 -a --output build/md5sum.txt.sig  --detach-sig build/md5sum.txt
	gpg2 --verify build/md5sum.txt.sig build/md5sum.txt
	find build/ -name "$(app)*" -exec sha1sum {} + > build/sha1sum.txt
	gpg2 -a --output build/sha1sum.txt.sig --detach-sig build/sha1sum.txt
	gpg2 --verify build/sha1sum.txt.sig build/sha1sum.txt
	@echo ""
	@echo ""
	@echo "MD5 hashes"
	@echo "========================"
	@cat build/md5sum.txt
	@echo ""
	@echo ""
	@echo "SHA1 hashes"
	@echo "========================"
	@cat build/sha1sum.txt
	@echo ""
	@echo ""
	@echo "*.sig files are signed with my GPG key"

clean:
	rm build/*.txt -f
	rm build/*.txt.sig -f
	rm build/telegramnotify -f
	rm build/telegramnotify.exe -f
	rm build/telegramnotify_macos -f
	rm build/telegramnotify*.rpm -f
	go clean

test: check

start:
	go run main.go
