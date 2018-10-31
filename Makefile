ci: setup all report

all: clean vet lint test

setup:
	go get -u golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

go-env:
ifndef GOPATH
	$(error GOPATH is not set)
endif

clean: go-env
	go clean

vet: go-env
	go vet `go list ./...`

lint: go-env
	golint `go list ./...`

test: go-env
	go test --cover -p 1 `go list ./...`

doc: go-env
	godoc --http=:6060

profile: go-env
	go test --cpuprofile cpu.prof --memprofile mem.prof -bench .

report:
	go test -v -covermode=count -coverprofile=coverage.out
	$(go env GOPATH | awk 'BEGIN{FS=":"} {print $1}')/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN

