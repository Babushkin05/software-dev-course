.PHONY: test-coverage

test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out -coverpkg=./... ./...
	@go tool cover -func=coverage/coverage.out | tee coverage/coverage.txt
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report generated:"
	@echo "- Text summary: coverage/coverage.txt"
	@echo "- HTML report: file://$(shell pwd)/coverage/coverage.html"