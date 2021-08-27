FROM golang:alpine
COPY ./../ /app/go
WORKDIR /app/go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

EXPOSE 80