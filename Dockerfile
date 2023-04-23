FROM registry.cn-hangzhou.aliyuncs.com/horizoncd/golang:1.19.8-alpine as build

WORKDIR /

COPY . /

RUN go build -o /hello-world

FROM registry.cn-hangzhou.aliyuncs.com/horizoncd/golang:1.19.8-alpine

WORKDIR /

COPY --from=build /hello-world /bin/hello-world

EXPOSE 8080

ENTRYPOINT [ "hello-world" ]