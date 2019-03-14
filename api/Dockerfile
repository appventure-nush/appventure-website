FROM golang:1.11-alpine as build

WORKDIR /go/src/github.com/appventure-nush/appventure-website/api
RUN apk add --no-cache git
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go get ./...
RUN go get github.com/ponzu-cms/ponzu/...
#RUN ponzu build -ldflags '-extldflags "-static"' -o /api
RUN ponzu build

ENTRYPOINT ["./ponzu-server"]
CMD ["run", "--bind=0.0.0.0"]
