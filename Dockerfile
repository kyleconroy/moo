FROM scratch

COPY gopath/bin/moo /moo

EXPOSE 8080

ENTRYPOINT ["/moo"]
