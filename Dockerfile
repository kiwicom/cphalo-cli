FROM alpine:latest
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*
COPY bin/cphalo /bin/
ENTRYPOINT [ "cphalo" ]
