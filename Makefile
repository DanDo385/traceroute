.PHONY: build run clean

build:
	@./scripts/build.sh

run:
	@./scripts/run.sh

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf ./bin
