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
	rm -f *.out *.out *.prof *.prof
	go mod tidy

vet:
	go vet ./...

lint:
	golint ./...

test:
	go test -v -covermode=count -coverprofile=coverage.out

doc:
	godoc --http=:6060

check: test
	go tool cover -html=coverage.out

profile:
	go test --cpuprofile cpu.prof --memprofile mem.prof -bench .

report:
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)


