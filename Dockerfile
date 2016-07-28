FROM alpine

EXPOSE 4100
EXPOSE 9292
EXPOSE 9324

COPY ./goaws_linux_amd64 /
COPY ./conf/goaws.yaml /conf/
ENTRYPOINT ["/goaws_linux_amd64"]
