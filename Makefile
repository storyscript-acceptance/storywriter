test:
	go vet
	ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race
	acceptance/omg/run.sh

build:
	go build

docker:
	docker build . -t williammartin/storywriter

.PHONY: test build docker
