# 此例子为项目中遇到的一个Makefile文件，这里将要把这个文件分析明白。先看懂，后会写。

# makefile文件描述了项目的编译规则，比如哪些文件需要编译、哪些文件不需要编译、哪些文件需要先编译、哪些文件需要后编译、哪些文件需要重建等。
# 编写了Makefile文件后，运行make命令就可以在该目录下寻找makefile文件，按照makefile中描述的规则自动编译，可以省略掉手动编译中的参数选项和命令，非常方便。
# makefile支持多线程并发操作，可以缩短项目的编译时间。

SHELL=/usr/bin/env bash

all: build
.PHONY: all

unexport GOFLAGS

GOVERSION:=$(shell go version | cut -d' ' -f 3 | sed 's/^go//' | awk -F. '{printf "%d%03d%03d", $$1, $$2, $$3}')
ifeq ($(shell expr $(GOVERSION) \< 1015005), 1)
$(warning Your Golang version is go$(shell expr $(GOVERSION) / 1000000).$(shell expr $(GOVERSION) % 1000000 / 1000).$(shell expr $(GOVERSION) % 1000))
$(error Update Golang to version to at least 1.15.5)
endif

# git modules that need to be loaded
MODULES:=

CLEAN:=
BINS:=

ldflags=-X=github.com/memoio/go-mefs-v2/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))+$(shell date "+%F.%T%Z")
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="$(ldflags)"

mefs: $(BUILD_DEPS)
	rm -f mefs
	go build $(GOFLAGS) -o mefs ./app/mefs

.PHONY: mefs
BINS+=mefs

keeper: $(BUILD_DEPS)
	rm -f mefs-keeper
	go build $(GOFLAGS) -o mefs-keeper ./app/keeper

.PHONY: mefs-keeper
BINS+=mefs-keeper 

user: $(BUILD_DEPS)
	rm -f mefs-user
	go build $(GOFLAGS) -o mefs-user ./app/user

.PHONY: mefs-user
BINS+=mefs-user

provider: $(BUILD_DEPS)
	rm -f mefs-provider
	go build $(GOFLAGS) -o mefs-provider ./app/provider

.PHONY: mefs-provider
BINS+=mefs-provider


build: mefs keeper user provider

.PHONY: build

clean:
	rm -rf $(BINS)
.PHONY: clean