PROJECT?=github.com/johananl/gophercon-eu-workshop
APP?=gophercon
PORT?=8000
INTERNAL_PORT?=4444

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ./bin/${APP}

build: clean
	#CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH}
	go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/

run: build
	PORT=${PORT} INTERNAL_PORT=${INTERNAL_PORT} ./bin/${APP}

test:
	go test -race ./...