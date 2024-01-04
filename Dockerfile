# Build stage

FROM golang:1.19-alpine3.17 as builder
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

#build go binary
RUN go build -o library-api cmd/main.go

# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

#Run stage

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/library-api .
# COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app./config/config.yaml ./config.yaml
# COPY db/migrations ./migrations
# COPY start.sh .



EXPOSE 8083

CMD [ "/app/library-api" ]

# ENTRYPOINT [ "/app/start.sh" ]