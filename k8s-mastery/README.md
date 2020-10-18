[https://github.com/rinormaloku/k8s-mastery/tree/envoy](https://github.com/rinormaloku/k8s-mastery/tree/envoy)

[Routing with the Envoy proxy in under 10 minutes](https://rinormaloku.com/routing-envoy-10-minutes/)

### The SA-Feedback service

Stores users feedback if the Sentiment Analysis was correct or not in a SQLite database. In a real app it would be used to train the Sentiment Analysis model, in our case we use it as an opportunity to showcase **Kubernetes Volumes**.

#### Setting up the service

**Prerequisite:** install `dotnet core 2.1` 

To run the app execute the command below (from the directory of sa-feedback)

```
$ dotnet run
```

### Deploy

```
docker-compose --compatibility up --build
```
Check that the application is up and running in http://localhost/. It is! That was quick and simple. Let’s get a quick overview of how this works.


### Envoy Edge Proxy an Overview

The Sentiment Analysis application is composed of three external facing services:

- **SA-Frontend** is routed to on three cases:
1. Base path: /
2. Static files: /static
3. Images Files: .png, .ico etc.
- **SA-WebApp** is routed to when the request path is /sentiment.
- **SA-Feedback** is routed to when the request path is /feedback.
To be able to target the clusters the envoy needs an address, this is dependent on the environment, and in our case, we are using Docker Compose which creates an entry for every service, in other words with the following config in [docker-compose.yaml](./docker-compose.yaml):

  sa-frontend:
    build: 
      context: ./sa-frontend
    image: rinormaloku/sentiment-analysis-frontend:feedback
    networks:
      - envoymesh
    expose:
      - "80"
This way the other services in the network can access sa-frontend using its name sa-frontend. Which is used in the envoy configuration in the next section.

Envoy Static Configuration
In Envoy nomenclature services are called clusters (i.e. there can be more than one service in the cluster), and below we display the configuration for sa-frontend cluster to forward requests to the address sa-frontend, defined in the file external-envoy.yaml:

``` yaml
  clusters:
  - name: sa-frontend
    connect_timeout: 0.25s
    type: strict_dns
    lb_policy: round_robin
    http_protocol_options: {}
    hosts:
    - socket_address:
        address: sa-frontend
        port_value: 80
```
With the cluster defined we need to route to them, which is specified in the filters section for the listener on port 0.0.0.0:80, as shown below:

``` yaml
- match:
    path: "/"
  route:
    cluster: sa-frontend
- match:
    prefix: "/static"
  route:
    cluster: sa-frontend
- match:
    regex: '^.*\.(ico|png|jpg)$'
  route:
    cluster: sa-frontend
```

And this is it! With this simple config we are:
1. Listening for requests from the downstream (can be any client in our case it’s the browser) in port 80.
2. In case of matching any of the options, we are forwarding the requests to the upstream, i.e instances of sa-frontend cluster.
3. Docker compose has added a DNS entry that forwards requests of sa-frontend to its container.

