# Copyright (c) 2020 DistAlchemist
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

PROJECT=mongongo
GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""	
  $(error Please set the environment variable GOPATH before running `make`)
endif

GO := GO111MODULE=on go 
GOBUILD := $(GO) build $(BUILD_FLAG) 
# -v: print the names of packages as they are compiled.
# -count n
#     Run each test and benchmark n times (default 1).
#     If -cpu is set, run n times for each GOMAXPROCS value.
#     Examples are always run once.
# https://golang.org/pkg/cmd/go/internal/test/:
# The rule for a match in the cache is that the run involves the same
# test binary and the flags on the command line come entirely from a
# restricted set of 'cacheable' test flags, defined as -cpu, -list,
# -parallel, -run, -short, and -v. If a run of go test has any test
# or non-test flags outside this set, the result is not cached. To
# disable test caching, use any test flag or argument other than the
# cacheable flags. The idiomatic way to disable test caching explicitly
# is to use -count=1.
GOTEST := $(GO) test -v --count=1 --parallel=1 -p=1

TEST_LDFLAGS := "" 

PACKAGE_LIST := go list ./... | grep -vE "cmd" 
PACKAGES := $$($(PACKAGE_LIST))

CURDIR := $(shell pwd) 
export PATH := $(CURDIR)/bin/:$(PATH) 

# Targets 
.PHONY: clean test dev cli mg-server 

default: cli mg-server

dev: default test 

test:
	@echo "Running test in native mode."
	@export TZ='Asia/Shanghai';\
	LOG_LEVEL=fatal $(GOTEST) -cover $(PACKAGES)

cli:
	$(GOBUILD) -o bin/cli cmd/cli/main.go 

mg-server:
	$(GOBUILD) -o bin/mg-server cmd/mgserver/main.go

ci: default 
	@echo "Checking formatting"
	@test -z "$$(gofmt -s -l $$(find . -name '*.go' -type f -print) | tee /dev/stderr)"
	@echo "Running Go vet"
	@go vet ./...

format:
	@gofmt -s -w `find . -name '*.go' -type f ! -path '*/_tools/*' -print`

clean:
	$(GO) clean -i ./...
	rm -rf ./bin 
