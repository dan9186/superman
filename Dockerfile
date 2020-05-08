FROM golang:1.14

ARG BUILD_PATH

ADD . $BUILD_PATH
WORKDIR $BUILD_PATH

RUN GOOS=linux \
    go build \
    -o /service .


FROM ubuntu

COPY --from=0 /service service
COPY --from=gomicro/probe /probe probe

EXPOSE 4567

CMD ["/service"]
