FROM golang:1.18-alpine AS builder

RUN apk --update add --no-cache git

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct 

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY ./ .

RUN go build -ldflags "-s -w -extldflags '-static'" -o gopher

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
  && apk --update add --no-cache tzdata \
  && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \ 
  && echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /build/gopher /usr/bin/gopher
RUN chmod +x /usr/bin/gopher

ENTRYPOINT [ "/usr/bin/gopher" ]
