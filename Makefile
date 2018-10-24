.DEFAULT_GOAL := get

export GOPATH := $(abspath .)
export GOBIN  := $(GOPATH)/bin/$(shell uname -s | tr A-Z a-z)
export PATH   := $(GOBIN):$(PATH)

install:
	@go get -u github.com/mitchellh/gox
	@go get -u github.com/kardianos/govendor
	@go get -u github.com/tmthrgd/go-bindata/...
.PHONY: install

govendor-list:
	@cd src/hub && $(GOBIN)/govendor list
.PHONY: govendor-list

govendor: govendor-list
	@cd src/hub && $(GOBIN)/govendor sync
.PHONY: govendor

govendor-add: govendor-list
	@cd src/hub && $(GOBIN)/govendor add +e
.PHONY: govendor-add

version:
	@sed -e s/'\$$version'/"git $(shell git rev-parse HEAD | cut -c-7) built on $(shell date +"%Y.%m.%d %H:%M %Z")"/ < \
		src/hub/util/version.go.template > src/hub/util/version.go
.PHONY: version

compile: govendor version
	@$(GOBIN)/gox -rebuild -tags git \
		-osarch="darwin/amd64 linux/amd64" \
		-output=$(GOPATH)/bin/{{.OS}}/{{.Dir}} \
		hub/...
.PHONY: compile

get: version
	@go get -tags git hub
.PHONY: get

bindata:
	$(GOBIN)/go-bindata -o src/hub/bindata/bindata.go -pkg bindata \
		meta/hub-well-known-parameters.yaml \
		src/hub/api/requests/*.template \
		src/hub/initialize/hub.yaml.template \
		src/hub/initialize/hub-component.yaml.template
.PHONY: bindata

fmt:
	@go fmt hub hub/api hub/aws hub/cmd hub/compose hub/config hub/git hub/initialize hub/kube \
		hub/lifecycle hub/manifest hub/parameters hub/state hub/storage hub/util
.PHONY: fmt

# go get -u github.com/hhatto/gocloc/cmd/gocloc
loc:
	@gocloc src/hub --not-match-d='src/hub/(vendor|bindata)'
.PHONY: loc

clean:
	-@rm -rf .cache pkg bin/darwin bin/linux \
		src/github.com src/golang.org src/gopkg.in \
		src/hub/vendor/github.com src/hub/vendor/golang.org src/hub/vendor/gopkg.in
	-@find src -not -path "*src/hub*" -not -path "src" -type d -maxdepth 1 | xargs rm -rf
.PHONY: clean
