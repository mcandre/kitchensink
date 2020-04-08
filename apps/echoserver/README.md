# echoserver: a mirror unto HTTP request bodies

# EXAMPLE

```console
$ mkdir -p bin
$ go build -o bin/echoserver .
$ ./bin/echoserver
$ curl http://localhost:8080/ -d batmobile
batmobile
```

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.14+

# RUNTIME REQUIREMENTS

(none)
