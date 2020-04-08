# kitchensink: Squeezing as many software stacks as possible into scale-down container deployments

# EXAMPLE (MINIKUBE)

```console
$ minikube start
$ eval "$(minikube docker-env)"
$ sh -c "cd apps/echoserver && ./build.sh"
$ kubectl apply -f kubernetes/echoserver.yml
$ kubectl port-forward echoserver-56dd7f9db7-94k84 8080:8080
$ curl http://localhost:8080/ -d batmobile
batmobile
```

# EXAMPLE (DOCKER-COMPOSE)

```console
$ cd apps/echoserver
$ docker-compose up
$ curl http://localhost:8080/ -d batmobile
batmobile
```

# EXAMPLE (DOCKER)

```console
$ cd apps/echoserver
$ ./build.sh
$ docker run --rm -it -p 8080:8080 mcandre/echoserver:latest
$ curl http://localhost:8080/ -d batmobile
batmobile
```

# EXAMPLE (HOST)

```console
$ mkdir -p bin
$ go build -o bin/echoserver .
$ ./bin/echoserver
$ curl http://localhost:8080/ -d batmobile
batmobile
```

Note: curl offers a handy `-w "\n"` flag for enhanced shell prompt trailing.

# ABOUT

In the beginning were *containers*. And they were good. They would scale UP production to tremendous heights, 10x, 100x, 1Mx. What an electric bill! Alas, no one ever thought to scale back DOWN in development. Whoo can run the whole system? No one! Bugs go unsquashed. Much fire. Features are sidelined. Thus, progress sputters and halts. The dream fades...

That's where kitchensink comes in. Your application matters. When you have a problem, you want the only technical difference between your laptop and prod to be credentials and a replication factor. Anyone can offer "microservices". We offer *minimal deployments*, designed to squeeze complex systems onto your test rig, so that you have more breathing room to prototype, hotfix, and streamline your SDLC.

The DevOps world is going bonkers right now with expertise all over the map. Some of you may be subject matter experts. If you've already syscall-gated WASM CloudABI bundles building cross-platform with fully OS-agnostic orchestration DSL's, great! And some of you could be newbies practicing your first coding snippets. The wonder of microservices is that all kinds of individuals are working on this cool stuff. For everything in-between, *containers* bae.

# REQUIREMENTS

* a modern POSIXy shell, e.g. sh, bash, zsh, ksh, [Git Bash](https://git-scm.com/)
* [direnv](https://direnv.net/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) with memory: 4096, cpus: 4
* [curl](https://curl.haxx.se/)

## Recommended

* [Docker Desktop](https://www.docker.com/) (includes docker and docker-compose)

# WARNING FOR MINIKUBE - DOCKER DESKTOP INTEGRATION

Running `eval "$(minikube docker-env)"` in your shell session will integrate Docker CLI with minikube. This enables some simple convenience commands for minikube development such as `docker build` and `docker images`.

HOWEVER, minikube containers should always be deployed with `kubectl`. In the event that a container is deployed with raw `docker run` commands, then the minikube cluster can become corrupt.

https://github.com/kubernetes/minikube/issues/7497

If this happens, you may invoke `minikube delete` to get back to a functioning state.

Remember to unset all `minikube docker-env` environment variables prior to launching any raw `docker run` deployments.
