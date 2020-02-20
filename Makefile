## Copyright 2020 Tokenomy Technologies Pte. Ltd. All rights reserved.
## Use of this source code is governed by a MIT-style license that can be
## found in the LICENSE file.

.PHONY: all lint test

LINTER_OPTS := --enable-all \
	--disable gocyclo \
	--disable dupl \
	--disable maligned \
	--disable funlen \
	--disable godox \
	--disable gocognit \
	--disable wsl \
	--disable gomnd

all: lint test

lint:
	-golangci-lint run $(LINTER_OPTS) ./...

test:
	go test ./...