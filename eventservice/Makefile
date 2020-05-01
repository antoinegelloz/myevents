override PWD=$(shell pwd)

# Go parameters
override GOCMD=go
override GOBUILD=$(GOCMD) build
override GOCLEAN=$(GOCMD) clean
override GOTEST=$(GOCMD) test
override GOGET=$(GOCMD) get
override GOENV=$(GOCMD) env
override GOPRIVATE=github.com/agelloz/myevents

#external services (see docker compose file)
override MONGODB=mongodb

#user overridable variables
BINARY_NAME=eventservice
TEST_RESULTS=/tmp/$(BINARY_NAME)/test-results
COVER_RESULTS=/tmp/$(BINARY_NAME)/test-coverage
PKG=./...
SHORT=true
FAILFAST=false
REPORT=true
TIMEOUT=60s
RUN=".*"

#setting environment
export PORT=8080
export MONGO_HOST=mongodb://localhost:27017
export GIN_MODE=release
export HTTP_ORIGIN="*"

all: fix-fmt code-checks test

install:
	$(GOGET) \
		golang.org/x/lint/golint \
		github.com/vektra/mockery/.../ \
		github.com/jstemmer/go-junit-report \
		honnef.co/go/tools/cmd/staticcheck \
		github.com/client9/misspell/cmd/misspell \
		github.com/radovskyb/watcher/... \
		github.com/devimteam/microgen/cmd/microgen \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \
		$(PWD)/...

#external services rules
$(MONGODB):
	$(DCU) -d $(MONGODB)

#test/bench rules
test: $(MONGODB)
	@mkdir -p $(COVER_RESULTS) $(TEST_RESULTS);
	@TEST_RESULTS=$(TEST_RESULTS) \
		COVER_RESULTS=$(COVER_RESULTS) \
		SHORT=$(SHORT) \
		PKG=$(PKG) \
		FAILFAST=$(FAILFAST) \
		REPORT=$(REPORT) \
		TIMEOUT=$(TIMEOUT) \
		RUN=$(RUN) \
		bash -c $(PWD)/scripts/test.sh

bench: $(MONGODB)
	@go test $(PKG) -run=XXX -bench=.

#run rule
run: $(MONGODB)
	watcher -dotfiles=false -cmd="go run ." -startcmd

#fmt rules
fix-fmt:
	@bash -c $(PWD)/scripts/fix-fmt.sh

code-checks:
	@bash -c $(PWD)/scripts/code-checks.sh

#clean rules
clean:
	$(GOCLEAN)
	@rm -f $(PWD)/$(BINARY_NAME)
	@$(DCD) -v

clean-go-modcache:
	$(GOCLEAN) --modcache

clean-test-results:
	rm -rf $(COVER_RESULTS) $(TEST_RESULTS)

clean-all: clean clean-go-modcache clean-test-results
