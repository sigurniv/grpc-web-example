FROM golang:1.10 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/sigurniv/grpc-web-golang/service/game
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/improbable-eng/grpc-web/go/grpcweb
RUN go get -u google.golang.org/grpc
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
COPY ./static ./static

EXPOSE 50053

ENTRYPOINT ["./app"]