alL: test


.PHONY: build
build:
	@echo "Nothing to do."


.PHONY: test
test:
	go get -t ./...
	go test -v ./...


.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	go get -t ./...
	goconvey -cover -port=9042 -workDir="$(realpath .)" -depth=1


.PHONY: cover
cover:
	go get -t ./...
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out .
