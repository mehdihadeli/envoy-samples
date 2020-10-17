### Guid To Install Envoy Locally

[How to install Envoy Proxy on Ubuntu 18](https://www.liquidweb.com/kb/how-to-install-envoy-proxy-on-ubuntu-18/)

[Using the Envoy Docker Image](https://www.envoyproxy.io/docs/envoy/latest/start/start#using-the-envoy-docker-image)

[Sandboxes](https://www.envoyproxy.io/docs/envoy/latest/start/start#sandboxes)

# Envoy proxy examples and experiments

To run the samples in this project please install the following dependencies:
* Shipyard https://shipyard.run
* Docker https://docker.io
* Xdg https://www.howtoinstall.me/ubuntu/18-04/xdg-utils/

# Simple TCP Load Balancing [./tcp-loadbalancing](./tcp-loadbalancing)
This example shows how to loadbalance two Docker containers using Envoy

# Simple Routing [./routing-simple](./routing-simple)
This example shows how to route to two different containers using HTTP path

# Simple Routing Kubernetes [./routing-simple-k8s](./routing-simple-k8s)
This example shows how to route to two different containers using HTTP path running in Kubernetes

# WASM HTTP Filters for Consul Service Mesh [./wasm-filters](./wasm-filters)
This example shows how WASM HTTP filters can be used with Envoy proxy

## Install Shipyard

```
curl https://shipyard.run/install | bash -s
```

## Create the environment

```
shipyard run ./wasm-filters
```
