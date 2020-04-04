FROM dev-reg.wanxingrowth.com/shoppingmall/builder:20200304103054 AS builder

COPY . /app
RUN cd /app && make build_service_cross_only

FROM alpine

ENV LOG_LEVEL debug
ENV CONFIG_FILE /app/config.json
ENV SERVICE_ID 0

COPY --from=builder /app/builds/release/service /app/service
COPY ./config/example.json /app/config.json
COPY ./builds/docker/entrypoint.sh /app/entrypoint.sh
COPY ./builds/docker/timezone /etc/localtime

WORKDIR /app
ENTRYPOINT ["/app/entrypoint.sh"]