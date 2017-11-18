FROM scratch

COPY gopath/bin/moo /moo

ENTRYPOINT ["/moo"]
