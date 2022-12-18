FROM golang:alpine as builder

WORKDIR /beta_build

COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download
RUN go mod verify

# special thanks to https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -v -o beta .

FROM scratch

COPY --from=builder /beta_build/beta /beta

ENTRYPOINT ["/beta"]