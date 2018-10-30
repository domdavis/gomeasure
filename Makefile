ci: setup all

all: clean vet lint test

setup:
	go get -u golang.org/x/lint/golint

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


