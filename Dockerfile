FROM golang:1.11.0-alpine3.8 as builder

RUN apk update && apk upgrade && apk --no-cache add git

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080

CMD ["/go/bin/app"]
