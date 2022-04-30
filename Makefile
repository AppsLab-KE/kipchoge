.PHONY: test

test:
	go vet ./...

release:
ifdef VERSION
	@echo " Release version "
	@git tag ${VERSION}
endif
