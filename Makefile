VERSION=0.1.0
APP = shortlink
IMAGE_NAME = therealplato/$(APP):$(VERSION)

GOOS = $(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH = amd64

all: image

clean:
	rm -f ${ARTIFACT}

build: clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(APP) 

image: TAG ?= latest
image: GOOS = linux
image: build
	docker build -t $(IMAGE_NAME) .

image-and-push: TAG ?= latest
image-and-push: test image
	docker push $(IMAGE_NAME)

test:
	docker-compose run --rm test

run: GOOS = linux
run: build
	docker-compose up --build shortlink
