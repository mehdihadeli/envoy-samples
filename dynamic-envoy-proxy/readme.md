[Dynamic Envoy Proxy on Linux Machine](https://medium.com/cstech/dynamic-envoy-proxy-on-linux-machine-25ccf8b159be)

### Envoy Configuration Methods

There is two configuration method we have. One is the static configuration and the other one is dynamic configuration.
Here is the simple static configuration example has one port listen to 10000 and redirect all requests to google.come that comes from the port.

``` yaml
static_resources:
  listeners:
  - name: listener_0
    address:
      socket_address:
        protocol: TCP
        address: 0.0.0.0
        port_value: 10000
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
          stat_prefix: ingress_http
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["*"]
              routes:
              - match:
                  prefix: "/"
                route:
                  host_rewrite_literal: www.google.com
                  cluster: service_google
          http_filters:
          - name: envoy.filters.http.router
  clusters:
  - name: service_google
    connect_timeout: 30s
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
Save the file as envoy.yaml and easily run on a docker container with the command below.

``` bash
docker run --name=proxy-static -d \
    -p 84:10000 \
    -v $(pwd)/static-configuration/envoy/:/etc/envoy \
    -v $(pwd)/static-configuration/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
    envoyproxy/envoy:v1.16-latest
```

After that, all the requests to port `84` will be proxied to google.com

### Envoy Dynamic Configurations

The envoy.yaml file is important. The file is an entry point for Envoy. As seen on ExecStart command above. Envoy never tracking the changes on the file after it started. It is okay. There is no reason to change the file. It is really simple and does not contain any detail.
We will copy all the envoy configuration files to the folder /etc/envoy

**envoy.yaml**
``` yaml
node:
  id: id_1
  cluster: main
admin:
  access_log_path: /tmp/envoy_admin_access.log
  address:
    socket_address:
      address: 0.0.0.0
      port_value: "9901"
dynamic_resources:
  cds_config:
    path: /etc/envoy/cds.yaml
  lds_config:
    path: /etc/envoy/lds.yaml
```

The node section is required. Envoy exposes a local administration interface that can be used to query and modify different aspects of the server. There is two dynamic config file provided. One for cluster definitionsand another one for listeners.

**lds.yaml**

``` yaml
version_info: "0"
resources:
- '@type': type.googleapis.com/envoy.api.v2.Listener
  name: listener_0
  address:
    socket_address:
      address: 0.0.0.0
      port_value: "80"
  filter_chains:
  - filters:
    - name: envoy.http_connection_manager
      config:
        stat_prefix: ingress_http
        codec_type: AUTO
        server_name: default
        rds:
          route_config_name: local_route
          config_source:
            path: /etc/envoy/rds.yaml
        http_filters:
        - name: envoy.router
```

The listener_0 configuration listens to port 80 and redirects all the requests to the route definition(/etc/envoy/rds.yaml) named local_route.

**rds.yaml**

``` yaml
version_info: "0"
resources:
- '@type': type.googleapis.com/envoy.api.v2.RouteConfiguration
  name: local_route # route_config_name on the lds.yaml
  virtual_hosts:
  - name: "EnvoyNetCore"
    domains:
    - "envoynetcore.com"
    routes:
    - match:
        prefix: /
      route:
        cluster: "EnvoyNetCore" # cluster name on the cds.yaml we want to point to.
        timeout: 60s
```
The route definition that passes all the requests to the cluster named EnvoyNetCore. The cluster is known from the envoy.yaml file. (/etc/envoy/cds.yaml)

**cds.yaml**

``` yaml
version_info: "0"
resources:
- '@type': type.googleapis.com/envoy.api.v2.Cluster
  name: "EnvoyNetCore"
  connect_timeout: 5s
  lb_policy: ROUND_ROBIN
  type: EDS
  eds_cluster_config:
    service_name: "EnvoyNetCore"
    eds_config:
      path: /etc/envoy/eds.yaml
```

Here is the cluster definition. You can declare more than one cluster on there. Cluster definitions points to endpoint definitions(eds.yaml). Also, at this level, you can configure a circuit breaker for the cluster.

**eds.yaml**

``` yaml
version_info: "0"
resources:
- '@type': type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
  cluster_name: ""
  endpoints:
  - lb_endpoints:
    - endpoint:
        address:
          socket_address:
            address: 127.0.0.1 # The address of the application.
            port_value: "8181" # Listening port from the application.
```

The last part of the configuration file is eds.yaml. This file contains address and port values to call any endpoint we want to include for the cluster.

``` bash
docker run --name=proxy-dynamic -d \
    -p 85:10000 \
    -v $(pwd)/dynamic-configuration/envoy/:/etc/envoy \
    -v $(pwd)/dynamic-configuration/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
    envoyproxy/envoy:v1.16-latest
```

After that, all the requests to port `85` will be proxied to our `hello-world` containers

> You can define only one resource item at this point but more than one endpoints can be defined.

> You can change these configuration files on the fly but Envoy can’t apply the change until you move the file with the same location & same name. So, create the new configurations as separate files with different names and move the files with the original names to the same path/name of the original ones.

``` bash
mv /etc/envoy/eds.new.yaml /etc/envoy/eds.yaml
```