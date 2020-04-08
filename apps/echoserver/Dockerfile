FROM library/golang:1.14
COPY main.go /go
RUN go build -o /go/echoserver -ldflags "-linkmode external -extldflags -static" .

FROM scratch
COPY --from=0 /go/echoserver /bin/echoserver
EXPOSE 8080
ENTRYPOINT ["/bin/echoserver"]
