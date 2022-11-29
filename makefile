include .env
export
LDFLAGS="-X main.NotionAuthorizationSecret=${NOTION_TOKEN}"

run:
	go run -ldflags ${LDFLAGS} .

build:
	go build -ldflags ${LDFLAGS} .

install:
	go install .
