all: clean build vet lint test

ci: env all report

env:
ifndef COVERALLS_TOKEN
	$(error COVERALLS_TOKEN is not set)
endif

build:
	go build

clean:
	go clean
	rm -f c.out coverage.out cpu.prof mem.prof

vet:
	go vet ./...

lint:
	golint ./...

test:
	go test -coverprofile c.out ./...

doc:
	godoc --http=:6060

profile:
	go test --cpuprofile cpu.prof --memprofile mem.prof -bench .

report:
	go test -v -covermode=count -coverprofile=coverage.out
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)


