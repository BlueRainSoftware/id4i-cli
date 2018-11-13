FROM golang:1.11.2 as build-image
RUN mkdir -p $GOPATH/src/github.com/BlueRainSoftware/id4i-cli
ADD . $GOPATH/src/github.com/BlueRainSoftware/id4i-cli
WORKDIR $GOPATH/src/github.com/BlueRainSoftware/id4i-cli

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o id4i main.go
RUN cp id4i /


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-image /id4i /home/

WORKDIR /home
ENTRYPOINT ./id4i

CMD ["./id4i"]
