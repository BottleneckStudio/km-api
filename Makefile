GO ?= go

test: lint
	@echo '  -> running test'
	@$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo '  -> running test' @$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo
.PHONY: test

lint:
	@echo '  -> running golangci_lint'
	@$(go env GOPATH)/bin/golangci-lint run
	@echo
.PHONY: lint

dev:
	@air
.PHONY: dev

deploy: test up clean
	@echo "  -> done ✓"
.PHONY: deploy

up: up.json
	@echo "  -> deploying"
	@up
.PHONY: up

clean:
	#@rm -rf up.json
	@rm -rf ./dist/
.PHONY: clean

up.json:
	@echo "  -> creating up.json from template file(TODO)"
