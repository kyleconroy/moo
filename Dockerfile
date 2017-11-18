FROM scratch

COPY gopath/bin/cowsay /cowsay

ENTRYPOINT ["/cowsay"]
