# echoserver: a mirror unto HTTP request bodies

# EXAMPLE

```console
$ echoserver

$ curl http://localhost:8080/ -d batmobile
batmobile
```

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.14+

# COMPILE

```console
$ mkdir bin
$ go build -o bin/echoserver .
```
