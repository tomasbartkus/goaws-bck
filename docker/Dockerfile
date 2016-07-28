FROM alpine:3.4

RUN apk add --update bash libc6-compat libxml2 && rm -rf /var/cache/apk/*
RUN ln -s /lib /lib64
RUN mkdir -p /scripts
RUN mkdir -p /bin
RUN mkdir -p /conf

COPY ./bin/goaws /bin/goaws
COPY ./goaws.yaml /conf/

CMD ["./bin/goaws"]

EXPOSE 4100