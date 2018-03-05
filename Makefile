# ref: https://github.com/dtan4/s3url/blob/master/Makefile
NAME := go-echo-sqlboiler
VERSION := 1.0.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR := $(shell go list ./... | grep -v vendor)

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

.PHONY: update-deps
update-deps: dep
	dep ensure -update -v

.PHONY: dist
dist:
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: watch
watch:
ifeq ($(shell command -v looper 2> /dev/null),)
	go get -u github.com/nathany/looper
endif
	looper

.PHONY: swag
swag:
ifeq ($(shell command -v swag 2> /dev/null),)
	go get -u github.com/swaggo/swag/cmd/swag
endif
	swag i

.PHONY: run
run:
ifeq ($(shell command -v gin 2> /dev/null),)
	go get -u github.com/codegangsta/gin
endif
	gin -a 1313 -p 1314 -i --notifications run main.go

.PHONY: release
release:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: cross-build
cross-build: deps
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done
