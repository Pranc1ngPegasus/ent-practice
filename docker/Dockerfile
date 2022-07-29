FROM golang:1.18 as builder

WORKDIR /go/src/app

ENV CGO_ENALBED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download -x

COPY . .

RUN go build -v -o server ./cmd/server

WORKDIR /go/src/app

FROM gcr.io/distroless/static:latest-amd64 as runner

COPY --from=builder /go/src/app/server /server

ENTRYPOINT ["/server"]
