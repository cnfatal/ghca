FROM golang
RUN go get github.com/fatalc/ghca
EXPOSE 8080
ENTRYPOINT ["/go/bin/ghca"]
CMD ["--server"]