.PHONY: test
test:
	go get -t ./...
	go test -v ./...


.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	go get -t ./...
	goconvey -cover -port=9042 -workDir="$(realpath .)" -depth=1
