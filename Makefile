VERSION=0.1.0
APP = shortlink
IMAGE_NAME = therealplato/$(APP):$(VERSION)

GOOS = $(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH = amd64

all: image

clean:
	docker-compose kill
	rm -f ${ARTIFACT}

build: GOOS = linux
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
	go test

run: build
	docker-compose up -d --build shortlink postgres

migrate: run
  docker-compose exec postgres psql -Udocker -c "insert into shortlink (slug, link) values ('abc', 'http://therealplato.com')" docker

bake: migrate
	curl -s localhost:8000 > testdata/root.golden.html
	curl -s localhost:8000/preview/abc > testdata/preview.golden.html
