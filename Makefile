all: clean build vet lint test

ci: env setup all report

env:
ifndef COVERALLS_TOKEN
	$(error COVERALLS_TOKEN is not set)
endif

setup:
	go get -u golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

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


