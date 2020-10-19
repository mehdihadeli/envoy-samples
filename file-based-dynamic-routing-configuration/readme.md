### Envoy Dynamic Configuration

(dynamic configuration)[https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration]

In the previous scenarios we've defined the `static configuration`. However this has made it `difficult to reload the configuration` when `changes` are required. To solve this, the static configuration can be defined as Dynamic Configuration. With `Dynamic Configuration`, when `changes` are made, Envoy will `automatically reload the changes` and apply them to the configuration and traffic routing.

Envoy supports different parts of the configuration as dynamic. The APIs available are:

- **EDS**: The Endpoint Discovery Service (EDS) API provides a way Envoy can discover members of an upstream cluster. This allows you to dynamically add and remove servers handling the traffic.

- **CDS**: The Cluster Discovery Service (CDS) API layers on a mechanism by which Envoy can discover upstream clusters used during routing.

- **RDS**: The Route Discovery Service (RDS) API layers on a mechanism by which Envoy can discover the entire route configuration for an HTTP connection manager filter at runtime. This would enable concepts such as dynamically changing traffic shifting and blue/green releases.

- **LDS**: The Listener Discovery Service (LDS) layers on a mechanism by which Envoy can discover entire listeners at runtime.

- **SDS**: The Secret Discovery Service (SDS) layers on a mechanism by which Envoy can discover cryptographic secrets (certificate plus private key, TLS session ticket keys) for its listeners, as well as the configuration of peer certificate validation logic (trusted root certs, revocations, etc).

The value for configuration can come from the filesystem, REST-JSON or gRPC endpoints.

More information can be found in the [xDS configuration API overview](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration)

we'll change our configuration to use Endpoint Discovery Service (EDS) allowing nodes to be dynamically added based with data coming from the `filesystem`.

### Cluster ID

An initial outline of the Envoy configuration required is available at envoy.yaml

The first changes required is to add a [Node](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#core-node). This allows the Envoy node to be identified, potentially allowing for unique configurations to be applied.

Prepend the following snippet to the top of the envoy.yaml file.

```
node:
  id: id_1
  cluster: test
```
The API also has support for additional metadata, such as [locality](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#core-locality) for providing region and zone-based information.

### EDS Configuration

The EDS configuration is defined to allow the upstream clusters to be controlled dynamically.

Within the `static configuration`, this was defined as:

``` yaml
  clusters:
  - name: targetCluster
    connect_timeout: 0.25s
    type: STRICT_DNS
    dns_lookup_family: V4_ONLY
    lb_policy: ROUND_ROBIN
    hosts: [
      { socket_address: { address: 172.18.0.3, port_value: 80 }},
      { socket_address: { address: 172.18.0.4, port_value: 80 }}
    ]
```

### Convert to EDS
To convert this to EDS based a **eds_cluster_config** is required and changing the type to **EDS**.

Add the following cluster to the end of the Envoy configuration.

``` yaml
clusters:
  - name: targetCluster
    connect_timeout: 0.25s
    lb_policy: ROUND_ROBIN
    type: EDS
    eds_cluster_config:
      service_name: localservices
      eds_config:
        path: '/etc/envoy/eds.conf'
```

The values for the upstream servers, such as 172.18.0.3 and 172.18.0.4, will come from the file [eds.conf]().

### EDS Configuration

The contents of [eds.conf](./eds.conf) is a JSON definition of the same information defined within our static configuration.

Create [eds.conf](./eds.conf) file with the following content:

``` yaml
{
  "version_info": "0",
  "resources": [{
    "@type": "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment",
    "cluster_name": "localservices",
    "endpoints": [{
      "lb_endpoints": [{
        "endpoint": {
          "address": {
            "socket_address": {
              "address": "172.17.0.1", #hostname -I    -this is ip address of host for docker, localhost dosn't work because we are in another container
              "port_value": 8001
            }
          }
        }
      }]
    }]
  }]
}
```

[https://askubuntu.com/questions/430853/how-do-i-find-my-internal-ip-address](https://askubuntu.com/questions/430853/how-do-i-find-my-internal-ip-address)

we can use `hostname -I` to get host address for our docker container and inner our envoy configuration we can't use `host.docker.internal` because it just work in context docker like `dockerfile` and `docker-compose` but we can't use it inner our envoy configuration and we should use host address in order to access other docker container such as access to `containersol/hello-world` container on expose port `8001`. we can test it after get host ip with `hostname -I`

```
curl 172.17.0.1:8001
```

### Start Envoy

With the Envoy configuration and EDS conf defined, it's now possible to start the proxy.
Launch the container with the following command in root our project:

``` yaml
docker run --name=proxy-eds-filebased -d -p 9901:9901 -p 80:10000 -v $(pwd)/envoy/:/etc/envoy/  envoyproxy/envoy:v1.16-latest
```

Start two HTTP servers to handle the incoming requests.

```
docker run --name hello-world1 -d -p 8001:80 containersol/hello-world
docker run --name hello-world2 -d -p 8002:80 containersol/hello-world
```

Based on the current EDS configuration, Envoy will send all the traffic to a single node. Try it with `curl localhost`

