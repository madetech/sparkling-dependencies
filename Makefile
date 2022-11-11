.PHONY: test-watch
test-watch:
	./test-watch.sh

.PHONY: test
test: build-test
	docker run -t sparkling-dependencies-test go test ./...

.PHONY: build-test
build-test:
	docker build --target=test -t sparkling-dependencies-test .
