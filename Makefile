PROJECT?=github.com/johananl/gophercon-eu-workshop
APP?=gophercon
PORT?=8000
INTERNAL_PORT?=4444

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=docker.io/johananl/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ./bin/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/


container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

run: container
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && docker rm $(CONTAINER_IMAGE):$(RELEASE) || true
	docker run --rm --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} \
		$(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -race ./...