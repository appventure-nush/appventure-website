FROM node:10-alpine as grunt

WORKDIR /src
RUN apk add --no-cache make=4.2.1-r2 python2=2.7.15-r1
COPY . .
RUN yarn install
RUN node ./node_modules/grunt-cli/bin/grunt


FROM golang:1.11-alpine as build

WORKDIR /go/src/github.com/appventure-nush/appventure-website
RUN apk add --no-cache git=2.18.0-r0
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go get ./...
RUN go build -ldflags '-extldflags "-static"' -o /website


FROM scratch

EXPOSE 8080

COPY --from=build /website /website
COPY --from=grunt /src /

ENTRYPOINT ["/website"]
