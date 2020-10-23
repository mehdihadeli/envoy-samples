### Introduction

In the previous scenarios we've defined the `static configuration` and `dynamic configuration using files`. In this scenario we will cover another type of `dynamic configuration: API dynamic configuration`.

The endpoint discovery service is a xDS management server based on `gRPC` or `REST-JSON` API server used by Envoy to fetch cluster members. The cluster members are called endpoint in Envoy terminology. For each cluster, Envoy fetch the endpoints from the discovery service. EDS is the preferred service discovery mechanism for a few reasons:

- Envoy has explicit knowledge of each upstream host (vs. routing through a DNS resolved load balancer) and can make more intelligent load balancing decisions.

- Extra attributes carried in the discovery API response for each host inform Envoy of the host’s load balancing weight, canary status, zone, etc. These additional attributes are used globally by the Envoy mesh during load balancing, statistic gathering, etc.

The Envoy project provides reference gRPC implementations of EDS and other discovery services in both Java and Go.

we'll change our configuration to use `Endpoint Discovery Service (EDS)` allowing nodes to be dynamically added based with data coming from the REST-JSON API.

[EDS](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#eds)

[Endpoint discovery service (EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#endpoint-discovery-service-eds)

[Mostly static with dynamic EDS](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples#mostly-static-with-dynamic-eds)

[Dynamic](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples#dynamic)

[config.cluster.v3.Cluster.EdsClusterConfig](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html?highlight=eds#config-cluster-v3-cluster-edsclusterconfig)

[config.core.v3.ConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/config_source.proto#config-core-v3-configsource)

[config.core.v3.ApiConfigSource](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/config_source.proto#config-core-v3-apiconfigsource)

[Examples](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples.html?highlight=examples)

### EDS Configuration

An initial outline of the Envoy configuration required is available at [envoy.yaml](envoy/envoy.yaml).

The first change required is to add a cluster configuration, with type EDS, and indicate that in eds_config should be using the REST API:

``` yaml
clusters:
  - name: targetCluster
    type: EDS
    connect_timeout: 0.25s
    eds_cluster_config: #https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html?highlight=eds#config-cluster-v3-cluster-edsclusterconfig
      service_name: myservice
      eds_config: #https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/config_source.proto#envoy-v3-api-msg-config-core-v3-configsource
        api_config_source: #https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/config_source.proto#config-core-v3-apiconfigsource
          api_type: REST
          cluster_names: [eds_cluster] #Cluster names should be used only with REST.The cluster with name cluster_name must be statically defined and its type must not be EDS
          refresh_delay: 5s
```

Note: `api_type` is set to v2 REST endpoint. If you want to switch to v1 simply use `api_type: REST_LEGACY`

After that you need to define how `eds_cluster` are `resolved`. For this example we are gonna use an `static configuration`:

``` yaml
- name: eds_cluster
    type: STATIC
    connect_timeout: 0.25s
    hosts: [{ socket_address: { address: 127.0.0.1, port_value: 8080 }}]
```
in this example with dynamic endpoint discovery via an EDS REST management server listening on `172.18.0.4:8080` is provided above.

Notice above that `eds_cluster` is defined `to point Envoy at the management server` on specific ip address and port for our management server here (172.18.0.4:8080). Even in an otherwise completely dynamic configurations, some static resources need to be defined to point Envoy at its xDS management server(s).

Launch Envoy with the following command:

``` bash
docker run --name=api-eds -d \
    -p 9901:9901 \
    -p 80:10000 \
    -v $(pwd)/envoy/:/etc/envoy \
    envoyproxy/envoy:v1.16-latest
```

or

Get Envoy binary

```
docker cp docker create envoyproxy/envoy:v1.15.0:/usr/local/bin/envoy .
```

So start envoy with debug enabled:

``` bash
envoy -c envoy_config.yaml -l debug
```
[Copying files from Docker container to host](https://stackoverflow.com/questions/22049212/copying-files-from-docker-container-to-host)

### Start upstream services

Now you have to start the upstream cluster. For this we are gonna use one example application:

```
docker run --name hello-world1 -d -p 8001:80 containersol/hello-world
docker run --name hello-world2 -d -p 8002:80 containersol/hello-world
```

You could test your upstream service executing the following command: 

``` bash
curl http://localhost:8001
curl http://localhost:8002
```

At this point we have the `Envoy started`, and the `upstream cluster started`, but they are `not connected` yet because the `eds_cluster` that we specified in the configuration is not started yet(its xds server is not running).

```
docker logs api-eds 
```
If you inspect the logs of Envoy, you should see errors when Envoy try to fetching the endpoints:


``` bash
[2020-10-22 18:30:24.006][8][warning][config] [source/common/config/http_subscription_impl.cc:113] REST update for /v2/discovery:endpoints failed
```

we need to start a [eds server](https://github.com/salrashid123/envoy_discovery#start-sds):

we use this example for our eds server that implemented with python.

[https://github.com/salrashid123/envoy_discovery#start-sds](https://github.com/salrashid123/envoy_discovery#start-sds)

``` bash
cd eds_server/

virtualenv env --python=/usr/bin/python3.8
source env/bin/activate
pip install -r requirements.txt

# ImportError: No module named enum
# pip install enum34

python main.py
```
[Can't connect to docker container running inside WSL2](https://github.com/microsoft/WSL/issues/4983#issuecomment-602487077)

[Accessing resources with IPv4 Networks](https://shipyard.run/docs/wsl2/)

[can't access container in WSL from Windows](https://github.com/microsoft/WSL/issues/4170#issuecomment-502818570)

[How to access docker containers on wsl 2](https://stackoverflow.com/questions/62753270/how-to-access-docker-containers-on-wsl-2)

[How to access host ip and port?](https://github.com/Microsoft/WSL/issues/1032#issuecomment-535764958)

[Accessing Linux networking apps from Windows (localhost)](https://docs.microsoft.com/en-us/windows/wsl/compare-versions#accessing-linux-networking-apps-from-windows-localhost)

There is currently a bug with WSL2 and Docker bindings for localhost when accessed via the IPv4 IP address `127.0.0.1`. This means that services running in Docker on WSL2 can not be accessed via the ip address `127.0.0.1` from outside the WLS2 container (the hostname localhost functions correctly). In order for magic URLs to function on WSL2, requires IPv6 networking to be enabled. IPv6 is not affected by the Docker bind bug and resources can be accessed both internally and externally.

on windows this command shows that our containers ip and ports in windows listen on `ipv6 addresses`

``` powershell
 netstat -a

 TCP    [::1]:8000             DESKTOP-5OAUHA4:0      LISTENING
 TCP    [::1]:8001             DESKTOP-5OAUHA4:0      LISTENING
 TCP    [::1]:8002             DESKTOP-5OAUHA4:0      LISTENING
 TCP    [::1]:9000             DESKTOP-5OAUHA4:0      LISTENING
```

`WSL2` only binds to `ipv6` localhost ([::1]) and not ipv4 localhost (127.0.0.1).
I resorted to using the ipv6 localhost url, `http://[::1]:8001` and it just worked.

we can use `ip addr | grep eth0` to get host address for our docker container and inner our envoy configuration we can't use `host.docker.internal` because it just work in context docker like `dockerfile` and `docker-compose` but we can't use it inner our envoy configuration and we should use host address in order to access other docker container such as access to `containersol/hello-world` container on expose port `8001`. we can test it after get host ip with `ip addr | grep eth0`

``` bash
curl 172.27.211.139:8001

or

curl localhost:8001

or

curl [::1]:8001
```

```
curl http://localhost:8080/
or
curl 172.27.211.139:8080
```

we should see the following output on SDS stdout indicating an inbound Envoy discovery request:

```
Inbound v2 request for discovery.  POST payload: {u'node': {u'build_version': u'fd44fd6051f5d1de3b020d0e03685c24075ba388/1.6.0-dev/Clean/RELEASE', u'cluster': u'mycluster', u'id': u'test-id'}, u'resource_names': [u'myservice']}
127.0.0.1 - - [29/Apr/2018 22:59:04] "POST /v2/discovery:endpoints HTTP/1.1" 200 -
```

then on the envoy proxy stdout, something like:

``` bash
[2018-04-29 22:59:10.323][157796][debug][config] bazel-out/k8-opt/bin/source/common/config/_virtual_includes/http_subscription_lib/common/config/http_subscription_impl.h:67] Sending REST request for /v2/discovery:endpoints
[2018-04-29 22:59:10.323][157796][debug][router] source/common/router/router.cc:250] [C0][S636378528925215024] cluster 'eds_cluster' match for URL '/v2/discovery:endpoints'
[2018-04-29 22:59:10.323][157796][debug][router] source/common/router/router.cc:298] [C0][S636378528925215024]   ':method':'POST'
[2018-04-29 22:59:10.323][157796][debug][router] source/common/router/router.cc:298] [C0][S636378528925215024]   ':path':'/v2/discovery:endpoints'
[2018-04-29 22:59:10.323][157796][debug][router] source/common/router/router.cc:298] [C0][S636378528925215024]   ':authority':'eds_cluster'
...
[2018-04-29 22:59:10.324][157796][debug][client] source/common/http/codec_client.cc:52] [C2] connected
[2018-04-29 22:59:10.324][157796][debug][pool] source/common/http/http1/conn_pool.cc:225] [C2] attaching to next request
...
[2018-04-29 22:59:10.330][157796][debug][client] source/common/http/codec_client.cc:81] [C2] response complete
[2018-04-29 22:59:10.330][157796][debug][pool] source/common/http/http1/conn_pool.cc:200] [C2] response complete
...
[2018-04-29 22:59:10.331][157796][debug][pool] source/common/http/http1/conn_pool.cc:115] [C2] client disconnected
```

Basically, this shows no updates were received from the endpoint

You can verify that envoy doesn't know anything about this endpoint by attempting to connect through to it (envoy listener):

```
curl -v  http://localhost:80
```

```
* Connection #0 to host localhost left intact
no healthy upstreams
```

### Add endpoint to EDS

Now we're ready to add an upstream service configuration to the EDS server.

connect to SDS servers UI console in browser at:
```
http://localhost:8080/
```

#### Create Endpoint

Since we defined the service as `myservice` in the `envoy_config.yaml`, we can need to register an endpoint against it:

``` json
curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{
  "hosts": [  //This field is deprecated. Set the load_assignment field instead.
    {
      "ip_address": "172.27.211.139", // in wsl we have to use wsl eth0 interface address or wsl ipv6 address or use localhost for address that it resolve to ipv6 address
      "port": 8001, //our upstream service
      "tags": {
        "az": "us-central1-a",
        "canary": false,
        "load_balancing_weight": 50
      }
    }
  ]
}' http://localhost:8080/edsservice/myservice
```
What this will do is set some endpoints for `myservice`. Now, envoy will query SDS for membership.

Verify the endpoint is registered:

``` bash
curl -X GET "http://localhost:8080/edsservice/myservice" -H "accept: application/json"
```

``` json
{
  "hosts": [ //This field is deprecated. Set the load_assignment field instead.
    {
      "ip_address": "[::1]", //// in wsl we have to use wsl eth0 interface address or wsl ipv6 address or use localhost for address that it resolve to ipv6 address
      "port": 8001, 
      "tags": {
        "az": "us-central1-a",
        "canary": false,
        "load_balancing_weight": 50
      }
    }
  ]
}
```
### Check client connectivity via envoy

Since we already started the upstream service above, you can connect to it via envoy:

``` bash
curl -i http://localhost
```

### Delete Endpoint

Ok, so now we've dynamically added in an endpoint...lets remove it by the SDS server's custom API and emptying out its hosts: []

``` bash
curl -X PUT --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{
  "hosts": [  ]
}' http://localhost:8080/edsservice/myservice
```

Now try the endpoint, you should see `no healthy upstream` message from envoy

``` bash
$ curl -v  http://localhost:10000/
```

### Rinse and repeat
Ok, you can continue to play with the endpoints by adding and removing new upstream services on different ports:

eg:

```
$ curl 172.27.211.139:8001
$ curl 172.27.211.139:8002
```
and then using the API to add hosts to the SDS server (use the PUT endpoint to do that)

``` json
curl -X PUT --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{
    "hosts": [
   {
      "ip_address":"172.27.211.139",
      "port":8001,
      "tags":{
         "az":"us-central1-a",
         "canary":false,
         "load_balancing_weight":50
      }
   },
   {
      "ip_address":"[::1]",
      "port":8002,
      "tags":{
         "az":"us-central1-a",
         "canary":false,
         "load_balancing_weight":50
      }
   }
]
    }' http://localhost:8080/edsservice/myservice
```

``` bash
curl -X GET "http://localhost:8080/edsservice/myservice" -H "accept: application/json"
```
Now we can verify that the traffic is balanced with all the nodes registered with the following command:

``` bash
while true; do curl http://localhost; sleep .5; printf '\n'; done
```
You will see different responses according to the service that processed the request.