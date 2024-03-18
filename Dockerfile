FROM golang:1.22 AS builder

WORKDIR /gocloudisk

COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o clouddisk \ 
    && chmod +x clouddisk

FROM alpine:latest

COPY --from=builder /gocloudisk/clouddisk /usr/bin/clouddisk

RUN chmod +x /usr/bin/clouddisk

ENTRYPOINT ["clouddisk"]