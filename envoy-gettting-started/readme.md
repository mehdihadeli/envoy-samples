### Create Proxy Config

Envoy is configured using a YAML definition file to control the proxy's behaviour. In this step, we're building a configuration using the Static Configuration API. This means that all the settings are pre-defined within the configuration.

Envoy also supports Dynamic Configuration. This allows the settings to be discovered via an external source.

[Quick Start to Run Simple Example](https://www.envoyproxy.io/docs/envoy/latest/start/start#quick-start-to-run-simple-example)

### Resources
The first line of the Envoy configuration defines the API configuration being used. In this case, we're configuring the Static API, so the first line should be static_resources. Copy the snippet to the editor.

```
static_resources:
```

### Listeners
The beginning of the configuration defines the Listeners. A Listener is the networking configuration, such as IP address and ports, that Envoy listens to for requests. Envoy runs inside of a Docker Container, so it needs to listen on the IP address `0.0.0.0`. In this case, Envoy will listen on port `10000`.

Below is the configuration to define this setup. Copy the snippet to the editor.

``` yaml
listeners:
  - name: listener_0
    address:
      socket_address: { address: 0.0.0.0, port_value: 10000 }
```

### Filter Chains and Filters

With Envoy listening for incoming traffic, the next stage is to define how to process the requests. Each Listener has a set of filters, and different Listeners can have a different set of filters.

In this example, we'll proxy all traffic to Google.com (thanks Google!). The result: We should be able to request the Envoy endpoint and see the Google homepage appear, without the URL changing.

Filtering is defined using filter_chains. The aim of each filter is to find a match on the incoming request, to match it to the target destination. Copy the snippet to the editor.

``` yaml
filter_chains:
  - filters:
    - name: envoy.filters.network.http_connection_manager
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
        stat_prefix: ingress_http
        codec_type: AUTO
        route_config:
          name: local_route
          virtual_hosts:
          - name: local_service
            domains: ["*"]
            routes:
            - match: { prefix: "/" }
              route: { host_rewrite_literal: www.google.com, cluster: service_google }
        http_filters:
        - name: envoy.filters.http.router
```

The filter is using envoy.http_connection_manager, a built-in filter designed for HTTP connections. The details are as follows:

- **stat_prefix**: The human-readable prefix to use when emitting statistics for the connection manager.

- **route_config**: The configuration for the route. If the virtual host matches, then the route is checked. In this example, the route_config matches all incoming HTTP requests, no matter the host domain requested.

- **routes**: If the URL prefix is matched then a set of route rules defines what should happen next. In this case "/" means match the root of the request

- **host_rewrite**: Change the inbound Host header for the HTTP request.

- **cluster**: The name of the cluster which will handle the request. The implementation is defined below.

- **http_filters**: The filter allows Envoy to adapt and modify the request as it is processed.

### Clusters

When a request matches a filter, the request is passed onto a cluster. The cluster shown below defines that the host is google.com running over HTTPS. If multiple hosts had been defined, then Envoy would perform a Round Robin strategy.

Copy the cluster implementation to complete the configuration:

``` yaml
clusters:
- name: service_google
  connect_timeout: 0.25s
  type: LOGICAL_DNS
  # Comment out the following line to test on v6 networks
  dns_lookup_family: V4_ONLY
  lb_policy: ROUND_ROBIN
  load_assignment:
    cluster_name: service_google
    endpoints:
    - lb_endpoints:
      - endpoint:
          address:
            socket_address:
              address: www.google.com
              port_value: 443
  transport_socket:
    name: envoy.transport_sockets.tls
    typed_config:
      "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
      sni: www.google.com
```

### Admin

Finally, an admin section is required. The admin section is explained in more detail in the following steps.

``` yaml
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address: { address: 0.0.0.0, port_value: 9901 }
```

This structure defines the boilerplate for Envoy Static Configuration. The Listener defines the ports and IP address for Envoy. The listener has a set of filters to match on the incoming requests. Once a request is matched, it will be forwarded to a cluster.

You can view the full configuration on [Github](https://github.com/envoyproxy/envoy/blob/35887594c78e14ad021a8a074914bdfa907304cf/configs/google_com_proxy.yaml).


### Start Proxy

With the configuration in place, the container can be launched and provided to Envoy. With Docker, this is done via a Volume Mount to /etc/envoy/envoy.yaml.

### Start Envoy

```
 sudo mkdir $(pwd)/envoy
 sudo nano $(pwd)/envoy/envoy.yaml
```
Launch the Proxy, bound to port 80 with:

``` bash
docker run --name=proxy -d \
  -p 80:10000 \
  -v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
  envoyproxy/envoy:v1.16-latest
```

### View Envoy

Once started, you should be able to send HTTP requests to port 80 with `curl localhost`.

You can view this via your local browser with the URL [http://localhost/](http://localhost/) As you will see, the request is proxied to Google.com and you should see the Google homepage without the URL changing.


### Admin View

Envoy provides an administration view, allowing you to view configuration, stats, logs and other internal Envoy data.

The admin can be defined by adding an additional resource definition, where the port for the admin view is defined. The port should not conflict with other Listener configurations.

``` yaml
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address: { address: 0.0.0.0, port_value: 9901 }
```

### Start Admin

This Docker Container also exposes the admin port to the outside world. The resource configuration above will make the admin view available to the public and should be used for demonstration purposes only, see the [documentation](https://www.envoyproxy.io/docs/envoy/latest/operations/admin) on how to secure the admin portal.

To expose the admin portal, run the following command:

``` bash
docker run --name=proxy-with-admin -d \
    -p 9901:9901 \
    -p 10000:10000 \
    -v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
    envoyproxy/envoy:v1.16-latest
```

The dashboard is now available at [http://localhost:9901](http://localhost:9901)

> "The administration interface in its current form both allows destructive operations to be performed (e.g., shutting down the server) as well as potentially exposes private information (e.g., stats, cluster names, cert info, etc.). It is critical that access to the administration interface is only allowed via a secure network" [Envoy Documentation](https://www.envoyproxy.io/docs/envoy/latest/operations/admin)


### Route to Docker Containers (front-proxy-sample)

[Front Proxy](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy)

[https://github.com/envoyproxy/envoy/tree/master/examples/front-proxy](https://github.com/envoyproxy/envoy/tree/master/examples/front-proxy)

The final example uses Envoy to proxy traffic to different Python services based on the requested URL path.

### Configuration

The configuration of the application is defined as a Docker Compose file. We use a Docker Compose file because we want to run several containers simultaneously - one for the proxy and one for each of the individual services.

You can view the file by clicking [front-proxy/docker-compose.yml](./front-proxy/docker-compose.yml).

### Application

The service is a Python web application it also uses Envoy within the container to forward traffic to the Python application. It’s not necessary to have Envoy in front of the application.

[front-proxy/service.py](./front-proxy/service.py)

### Envoy Frontend Proxy

The Envoy proxy configuration is defined in: [samples/front-proxy/front-envoy.yaml](./samples/front-proxy/front-envoy.yaml)

As described in the first step, the configuration starts by defining a set of static_resources. The routes match based on the URL of the request.

``` yaml
routes:
    - match:
        prefix: "/service/1"
    route:
        cluster: service1
    - match:
        prefix: "/service/2"
    route:
        cluster: service2
```

The cluster configuration forwards traffic to endpoints called service1 and service2. These are DNS entries provided by the Docker Network, configured with Docker Compose.

### Deploy

Start the example using the Docker Compose command below:

``` bash
$ pwd
/home/mehdi
$ docker-compose -f samples/front-proxy/docker-compose.yaml build --pull
$ docker-compose -f samples/front-proxy/docker-compose.yaml up -d
$ docker-compose ps
```


### Test Envoy’s routing capabilities

You can now send a request to both services via the front-envoy.

For service1:

``` 
curl -v localhost:8080/service/1
```

For service2:

``` 
curl -v localhost:8080/service/2
```

Notice that each request, while sent to the front Envoy, was correctly routed to the respective application.

We can also use HTTPS to call services behind the front Envoy. For example, calling service1:

```
curl https://localhost:8443/service/1 -k -v
```

### Test Envoy’s load balancing capabilities

Now let’s scale up our service1 nodes to demonstrate the load balancing abilities of Envoy:

```
docker-compose scale service1=3
```

Now if we send a request to service1 multiple times, the front Envoy will load balance the requests by doing a round robin of the three service1 machines:

```
curl -v localhost:8080/service/1
curl -v localhost:8080/service/1
```

### enter containers and curl services

In addition of using curl from your host machine, you can also enter the containers themselves and curl from inside them. To enter a container you can use docker-compose exec <container_name> /bin/bash. For example we can enter the front-envoy container, and curl for services locally:

``` bash
docker-compose exec front-envoy /bin/bash

curl localhost:8080/service/1
curl localhost:8080/service/1
curl localhost:8080/service/2
```