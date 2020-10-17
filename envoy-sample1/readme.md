

### Guid To Install Envoy Locally

[How to install Envoy Proxy on Ubuntu 18](https://www.liquidweb.com/kb/how-to-install-envoy-proxy-on-ubuntu-18/)

[Using the Envoy Docker Image](https://www.envoyproxy.io/docs/envoy/latest/start/start#using-the-envoy-docker-image)

[Sandboxes](https://www.envoyproxy.io/docs/envoy/latest/start/start#sandboxes)

### Samples

(Envoy Proxy Crash Course, Architecture, L7 & L4 Proxying, HTTP/2, Enabling TLS 1.2/1.3 and more)[https://www.youtube.com/watch?v=40gKzHQWgP0]

Use this command for running needed container for our apps

we use [shipyard](https://github.com/shipyard-run/shipyard) for managing our container it like docker compose and therefrom

```
shipyard run ./
```

for destroy our container

```
shipyard destroy
```

for set envoy configuration with each of oour configurations we can use a command like this

[http.yaml](./http.yaml)
```
envoy --config-path http.yaml 
curl localhost:80
```

[allbackend.yaml](./allbackend.yaml)
```
envoy --config-path allbackend.yaml
curl localhost:8080
```

[app1app2split.yaml](./app1app2split.yaml)

```
envoy --config-path app1app2split.yaml
curl localhost:8080
```


[app1app2split.yaml](./blockadmin.yaml)

```
envoy --config-path blockadmin.yaml
curl localhost:8080
```


[tcp.yaml](./tcp.yaml)

```
envoy --config-path tcp.yaml 
curl localhost:8080
```