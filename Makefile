alL: test


.PHONY: build
build:
	@echo "Nothing to do."


.PHONY: test
test:
	go test -v .


.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=9042 -workDir="$(realpath .)" -depth=1


.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out .
