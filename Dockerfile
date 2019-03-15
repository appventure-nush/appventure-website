FROM node:10-alpine as grunt

WORKDIR /src
RUN apk add --no-cache make python2
COPY . .
RUN yarn install
RUN node ./node_modules/grunt-cli/bin/grunt

RUN yarn && yarn grunt
RUN yarn fmt

FROM golang:1.11-alpine as build

WORKDIR /go/src/github.com/appventure-nush/appventure-website
RUN apk add --no-cache git
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
