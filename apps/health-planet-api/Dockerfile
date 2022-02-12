FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o health-planet-api main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/health-planet-api /
COPY --from=buildImage /usr/src/app/config.localhost.yaml /
ENTRYPOINT ["/health-planet-api"]
