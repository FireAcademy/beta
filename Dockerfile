FROM golang:alpine as builder

WORKDIR /beta_build

COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download
RUN go mod verify

RUN go build -v -o beta .

FROM alpine

COPY --from=builder /beta_build/beta /usr/local/bin/beta

CMD beta