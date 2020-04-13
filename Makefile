GO ?= go

test: lint
	@echo '  -> running test'
	@$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo '  -> running test' @$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo
.PHONY: test

lint:
	@echo '  -> running golangci_lint'
	@golangci-lint run
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
	@rm -rf up.json
	@rm -rf ./dist/
.PHONY: clean


# parse up template
up.json:
	@echo "  -> creating up.json from template file"
	@cat up.tmpl | sed "s/\$$COGNITO_CLIENT_ID/${COGNITO_CLIENT_ID}/g" \
		| sed "s/\$$COGNITO_CLIENT_SECRET/${COGNITO_CLIENT_SECRET}/g" \
		| sed "s/\$$COGNITO_POOL_ID/${COGNITO_POOL_ID}/g" \
		| sed "s/\$$GITHUB_CLIENT_SECRET/${GITHUB_CLIENT_SECRET}/g" \
		| sed "s/\$$GITHUB_CLIENT_ID/${GITHUB_CLIENT_ID}/g" \
		| sed "s/\$$SLACK_TOKEN/${SLACK_TOKEN}/g" \
		| sed "s/\$$SESSION_KEY/${SESSION_KEY}/g" \
		| sed "s/\$$CSRF_KEY/${CSRF_KEY}/g" \
		| sed "s|\$$GITHUB_CALLBACK|${GITHUB_CALLBACK}|g" \
		| sed "s/\$$DYNAMO_TABLE_POSTS/${DYNAMO_TABLE_POSTS}/g" \
		| sed "s/\$$DYNAMO_TABLE_LIKES/${DYNAMO_TABLE_LIKES}/g" \
		> up.json

