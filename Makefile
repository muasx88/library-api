BINARY=library-api

dev:
	go run ./cmd/main.go

build:
	@ printf "Building aplication... "

	@ go build \
 		-trimpath \
 		-o target/${BINARY} \
 		./cmd/

	@ echo "done"

.PHONY: dev build