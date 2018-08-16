# Copyright 2018 Gavin Chun Jin. All rights reserved.
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

.PHONY: fmt test gometa benchmark

test:
	go test -cpu=1,2 .

fmt:
	go fmt .

gometa:
	test -n "$(GOBIN)"  # $$GOBIN
	${GOBIN}/gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow ./...

benchmark:
	go test . -bench .
