all: clean build test vet lint

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

profile:
	go test --cpuprofile cpu.prof --memprofile mem.prof -bench .

report:
ifndef COVERALLS_TOKEN
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
else
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)
endif


