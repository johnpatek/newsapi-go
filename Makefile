UNIT_TEST_HEADER        = "****************************** UNIT TEST *******************************"
LINT_TEST_HEADER        = "****************************** LINT TEST *******************************"
CODE_COVERAGE_HEADER    = "**************************** CODE COVERAGE *****************************" 
.PHONY: all
all: test lint

.PHONY: test
test: unit cover

.PHONY: unit
unit:
	@echo $(UNIT_TEST_HEADER)
	go test -v -timeout 30s -coverprofile=coverage.out ./...

.PHONY: cover
cover:
	@echo $(CODE_COVERAGE_HEADER)
	go tool cover -func=coverage.out

.PHONY: lint
lint:
	@echo $(LINT_TEST_HEADER)
	@if [ ! -f bin/golangci-lint ]; then \
    	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b bin; \
	fi
	./bin/golangci-lint -v run ./...

.PHONY: clean
clean:
	rm -rf coverage.out