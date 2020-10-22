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
    eds_cluster_config: #https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html?highlight=eds_cluster_config#config-cluster-v3-cluster-edsclusterconfig
      service_name: localservices
      eds_config:
        path: '/etc/envoy/eds.conf'
```

The values for the upstream servers, such as 172.18.0.3 and 172.18.0.4, will come from the file [eds.conf]().

### EDS Configuration

The contents of [eds.conf](envoy/eds.conf) is a JSON definition of the same information defined within our static configuration.

Create [eds.conf](envoy/eds.conf) file with the following content:

``` json
{
  "version_info": "0", #https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/discovery.proto#discoveryresponse
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

``` bash
curl 172.17.0.1:8001
```

### Start Envoy

With the Envoy configuration and EDS conf defined, it's now possible to start the proxy.
Launch the container with the following command in root our project:

``` bash
docker run --name=proxy-eds-filebased -d -p 9901:9901 -p 80:10000 -v $(pwd)/envoy/:/etc/envoy/  envoyproxy/envoy:v1.16-latest
```

Start two HTTP servers to handle the incoming requests.

```
docker run --name hello-world1 -d -p 8001:80 containersol/hello-world
docker run --name hello-world2 -d -p 8002:80 containersol/hello-world
```

Based on the current EDS configuration, Envoy will send all the traffic to a single node. Try it with `curl localhost`

### Endpoint Discovery Service (EDS) Configuration

[Endpoint Discovery Service (EDS) API ](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds)

[xDS API endpoints](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/xds_api#config-overview-management-server)

[Bootstrap configuration](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/bootstrap#config-overview-bootstrap)

[config.bootstrap.v3.Bootstrap](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/bootstrap/v3/bootstrap.proto#config-bootstrap-v3-bootstrap)

[config.cluster.v3.Cluster.EdsClusterConfig](https://github.com/envoyproxy/envoy/blob/5c7737266e671ea9801c14d2779ca30bb0032bf7/api/envoy/config/cluster/v3/cluster.proto#L180)

[Mostly static with dynamic EDS](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples)

[service.discovery.v3.DiscoveryResponse](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/discovery/v3/discovery.proto#service-discovery-v3-discoveryresponse)

[service.discovery.v3.DiscoveryRequest](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/discovery/v3/discovery.proto#service-discovery-v3-discoveryrequest)

[config.core.v3.ConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/config_source.proto#envoy-v3-api-msg-config-core-v3-configsource)

[supported config formats](https://www.envoyproxy.io/docs/envoy/latest/operations/cli#cmdoption-c)

[config.bootstrap.v2.Bootstrap.DynamicResources](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#config-bootstrap-v2-bootstrap-dynamicresources)

[core.ConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-configsource)

[core.ApiConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-apiconfigsource)

[xDS REST and gRPC protocol](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#xds-rest-and-grpc-protocol)

[Resource Types](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#resource-types)

[Filesystem subscriptions](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#filesystem-subscriptions)

With the configuration based on EDS, when the services need to be scaled up a new endpoint can be added to the `eds.conf`. Envoy will then automatically include the changes.

Replace the configuration with the following to add a new endpoint to the cluster.

``` json
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
              "address": "172.17.0.1",
              "port_value": 8001
            }
          }
        }
      },
        {
        "endpoint": {
          "address": {
            "socket_address": {
              "address": "172.17.0.1",
              "port_value": 8002
            }
          }
        }
      }]
    }]
  }]
}
```
Based on how Docker handles `file tracking` on our mapped volume, sometimes the filesystem change isn't triggered and detected. Force the change with the command `mv eds.conf tmp; mv tmp eds.conf`.

Envoy should automatically reload the configuration and add the new node into the load balancing rotation `curl localhost`.

### CDS Configuration

[Cluster discovery service](https://www.envoyproxy.io/docs/envoy/latest/configuration/upstream/cluster_manager/cds)

[CDS](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#cds)

[gRPC streaming endpoints](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/xds_api#grpc-streaming-endpoints)

[REST endpoints](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/xds_api#rest-endpoints)

[examples](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples)

[service.discovery.v3.DiscoveryResponse](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/discovery/v3/discovery.proto#service-discovery-v3-discoveryresponse)

[service.discovery.v3.DiscoveryRequest](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/discovery/v3/discovery.proto#service-discovery-v3-discoveryrequest)

[config.bootstrap.v2.Bootstrap.DynamicResources](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#config-bootstrap-v2-bootstrap-dynamicresources)

[core.ConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-configsource)

[core.ApiConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-apiconfigsource)

[xDS REST and gRPC protocol](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#xds-rest-and-grpc-protocol)

[Resource Types](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#resource-types)

[Filesystem subscriptions](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#filesystem-subscriptions)


With EDS in place, it's possible to move to scale up the upstream clusters. If we wanted to be able to dynamically add new domains and clusters, the Cluster Discovery Service (CDS) API needs to be implemented. In the following steps, we are configuring the Cluster Discovery Service (CDS) and The Listener Discovery Service (LDS).

You need to create a file to put the configuration for the clusters: [cds.conf](envoy/cds.conf).

bellow json file actually is a [DiscoveryResponse](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/discovery/v3/discovery.proto#envoy-v3-api-msg-service-discovery-v3-discoveryresponse)

``` json
{
  "version_info": "0",   
  "resources": [{
      "@type": "type.googleapis.com/envoy.api.v2.Cluster", #https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#resource-types
      "name": "targetCluster",
            "connect_timeout": "0.25s",
            "lb_policy": "ROUND_ROBIN",
            "type": "EDS",
            "eds_cluster_config": {
                "service_name": "localservices",
                "eds_config": {
                    "path": "/etc/envoy/eds.conf"
                }
            }
  }]
}
```

And also, you need to create a file to put the configuration for the listeners: [lds.conf](envoy/lds.conf).

[Listener discovery service (LDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/lds)

``` json
{
    "version_info": "0",
    "resources": [{
            "@type": "type.googleapis.com/envoy.api.v2.Listener",
            "name": "listener_0",
            "address": {
                "socket_address": {
                    "address": "0.0.0.0",
                    "port_value": 10000
                }
            },
            "filter_chains": [
                {
                    "filters": [
                        {
                            "name": "envoy.http_connection_manager",
                            "config": {
                                "stat_prefix": "ingress_http",
                                "codec_type": "AUTO",
                                "route_config": {
                                    "name": "local_route",
                                    "virtual_hosts": [
                                        {
                                            "name": "local_service",
                                            "domains": [
                                                "*"
                                            ],
                                            "routes": [
                                                {
                                                    "match": {
                                                        "prefix": "/"
                                                    },
                                                    "route": {
                                                        "cluster": "targetCluster"
                                                    }
                                                }
                                            ]
                                        }
                                    ]
                                },
                                "http_filters": [
                                    {
                                        "name": "envoy.router"
                                    }
                                ]
                            }
                        }
                    ]
                }
            ]
    }]
}
```

The content of files cds.conf and lds.conf is a `JSON definition` of with the same information defined within our static configuration but we can also use yaml.[Filesystem subscriptions](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#filesystem-subscriptions)

With the `externalized` the configuration of `clusters` and `listeners`, you need to `modify` your `Envoy's configuration` to make `reference` to these files. This can be accomplish changing all the `static_resources` for `dynamic_resources`.

Open the Envoy configuration file [envoy1.yaml](envoy/envoy1.yaml), and add the following configuration:

``` bash
dynamic_resources: #https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#config-bootstrap-v2-bootstrap-dynamicresources
  cds_config:
    path: "/etc/envoy/cds.conf" #https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-configsource
  lds_config:
    path: "/etc/envoy/lds.conf" #https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/config_source.proto#envoy-api-msg-core-configsource
```

After that, launch the container with the following command:

``` bash
docker run --name=proxy-eds-cds-lds-filebased -d \
-p 9902:9901 \
-p 81:10000 \
-v $(pwd)/envoy/:/etc/envoy \
-v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy1.yaml \
envoyproxy/envoy:v1.16-latest
```
Note: to avoid port conflicts, we exposed the ports with offset 1.

Execute the following command: `curl localhost:81`


### CDS Apply Changes

With the configuration based on CDS, LDS and EDS, we can dynamically add a new cluster.

Open the file [cds.conf](envoy/cds.conf).

### Add new cluster

We'll call this new cluster `newTargetCluster`. Replace the configuration with the following to add a new cluster.

``` json
{
  "version_info": "0",
  "resources": [{
      "@type": "type.googleapis.com/envoy.api.v2.Cluster",
      "name": "targetCluster",
            "connect_timeout": "0.25s",
            "lb_policy": "ROUND_ROBIN",
            "type": "EDS",
            "eds_cluster_config": {
                "service_name": "localservices",
                "eds_config": {
                    "path": "/etc/envoy/eds.conf"
                }
            }
  },
  {
      "@type": "type.googleapis.com/envoy.api.v2.Cluster",
      "name": "newTargetCluster",
            "connect_timeout": "0.25s",
            "lb_policy": "ROUND_ROBIN",
            "type": "EDS",
            "eds_cluster_config": {
                "service_name": "localservices",
                "eds_config": {
                    "path": "/etc/envoy/eds1.conf"
                }
            }
  }]
}
```

You also need to create the eds_cluster_config file for this new cluster. Create the file [eds1.conf](envoy/eds1.conf) with this content:

``` json
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
              "address": "172.17.0.1",
              "port_value": 8003
            }
          }
        }
      },
        {
        "endpoint": {
          "address": {
            "socket_address": {
              "address": "172.17.0.1",
              "port_value": 8004
            }
          }
        }
      }]
    }]
  }]
}
```
And you can use this new cluster, in the listener that you previously configured. Open the file lds.conf. Replace the target cluster with the name of the new cluster newTargetCluster.

``` json
"route": {
      "cluster": "newTargetCluster"
  }
```
The configuration of lds.conf should look like:

``` json
"filter_chains": [
    {
        "filters": [
            {
                "name": "envoy.http_connection_manager",
                "config": {
                    "stat_prefix": "ingress_http",
                    "codec_type": "AUTO",
                    "route_config": {
                        "name": "local_route",
                        "virtual_hosts": [
                            {
                                "name": "local_service",
                                "domains": [
                                    "*"
                                ],
                                "routes": [
                                    {
                                        "match": {
                                            "prefix": "/"
                                        },
                                        "route": {
                                            "cluster": "newTargetCluster"
                                        }
                                    }
                                ]
                            }
                        ]
                    },
                    "http_filters": [
                        {
                            "name": "envoy.router"
                        }
                    ]
                }
            }
        ]
    }
]
```

Start two other new HTTP servers to handle the incoming requests.

``` bash
docker run --name hello-world3 -d -p 8003:80 containersol/hello-world
docker run --name hello-world4 -d -p 8004:80 containersol/hello-world
```

Based on how Docker handles file inode tracking, sometimes the filesystem change isn't triggered and detected. Force the change with the command: 

``` bash
mv cds.conf tmp; mv tmp cds.conf; mv lds.conf tmp; mv tmp lds.conf
```

Envoy should automatically reload the configuration and add the new cluster. You can try running the following command: 

```
curl localhost:81
```

You can notice with the response of each request, that the ID of the nodes changes, corresponding to the nodes of newTargetCluster.