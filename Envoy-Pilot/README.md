# Envoy-Pilot 

[![Build Status](https://travis-ci.org/tak2siva/Envoy-Pilot.svg?branch=master)](https://travis-ci.org/tak2siva/Envoy-Pilot)    ![Version](https://img.shields.io/badge/version-v0.2.3-yellowgreen.svg)
![docker pulls](https://img.shields.io/docker/pulls/tak2siva/envoy-pilot)


Envoy Pilot or Envoy xDS is a control plane implementation for [Envoy](https://github.com/envoyproxy/envoy) written in Golang and uses Consul for `persistence` by default. It can also run without Consul by loading configuration from file.

This is an extension of [go-control-plane](https://github.com/envoyproxy/go-control-plane)

Currently Supports
   * CDS
   * LDS
   * RDS
   * EDS
   * ADS (for CDS & LDS)

*Note: Some infrequent configurations might not be mapped. Feel free to PR* 

Checkout [Envoy XDS PROTOCOL Overview](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol) for more detail

Also Checkout Sample Project 
   * [Envoy xDS Example From File](https://github.com/tak2siva/Envoy-xDS-Example-From-File)
   * [Envoy xDS Example With Consul](https://github.com/tak2siva/Envoy-xDS-Example-Consul)

## File Config
  Checkout the above example to load config from file

## Consul Usage

xDS Server will be exposed on port 7777

Run Envoy Proxy with the following configurations or use `--service-node` && `--service-cluster`. they do same things.

[core.Node](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#core-node)

``` yaml
node:
  id: ride-service-replica-2
  cluster: ride-service
```

``` bash
envoy -c ./envoy-config.yaml -l debug --service-cluster ride-service --service-node ride-service-replica-2

envoy -c ./envoy-config.yaml -l debug --service-cluster xdstest-cluster --service-node xdstest-node

or we can use node element in our envoy configs instead of --service-cluster, --service-node

envoy -c ./envoy-config.yaml -l debug 
```

our envoy config

``` yaml
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address: { address: 0.0.0.0, port_value: 9901 }
node:
  id: ride-service-replica-2
  cluster: ride-service
dynamic_resources:
  cds_config: {ads: {}}
  lds_config: {ads: {}}
  ads_config:
    api_type: GRPC
    grpc_services:
      envoy_grpc:
        cluster_name: xds_cluster
static_resources:
  clusters:
  - name: xds_cluster
    connect_timeout: 0.25s
    type: strict_dns
    lb_policy: ROUND_ROBIN
    dns_refresh_rate: 5s
    http2_protocol_options: {}
    hosts: [{ socket_address: { address: "localhost", port_value: 7777 }}]
```

[Aggregated xDS (“ADS”)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#aggregated-xds-ads)


[Aggregated Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/xds_api#aggregated-discovery-service)

some init scripts

``` bash
consul kv import @config/values.json

Consul KV Export

chmod +x config/wait-for-command.sh; sh config/wait-for-command.sh -t 30 -c \"curl -f http://localhost:8500/v1/status/leader\" && CONSUL_HTTP_ADDR=localhost:8500 consul kv import @config/values.json

apk add --no-cache curl; chmod +x config/wait-for-command.sh; sh config/wait-for-command.sh -t 30 -c \"curl -f http://localhost:8500/v1/status/leader\" && ./Envoy-Pilot
```

[Consul KV Import](https://www.consul.io/commands/kv/import)

[Consul KV Export](https://www.consul.io/commands/kv/export)


Every *DS requires two keys to be set in consul
  * config
  * version

And the key template is `xDS/app-cluster/ride-service/CDS/(config|version)`

For CDS add KV 

``` json
  * `xDS/app-cluster/ride-service/CDS/version` => "1.0"
  * `xDS/app-cluster/ride-service/CDS/config` => `"[
      {
        "name": "app_service",
        "connect_timeout": "0.250s",
        "type": "STRICT_DNS",
        "lb_policy": "RANDOM",
        "hosts": [{
          "socket_address": {
           "address": "localhost",
           "port_value": 8123
          }
        }]
    }
  ]"`
```

For LSD add KV 

``` json
  * `xDS/app-cluster/ride-service/LSD/version` => "1.0"
  * `xDS/app-cluster/ride-service/LSD/config` => `"
[
			{
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
															"cluster": "app_service"
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
			}
		]
  "`
```

#### Pushing new configuration

  * Envoy-Pilot will be polling for version change every 10 seconds.  
  * If there is a version mismatch for any of `xDS/app-cluster/ride-service/(CDS|LDS|RDS|EDS)/version` then new config `xDS/app-cluster/ride-service/(CDS|LDS|RDS|EDS)/config` will be pushed to subscriber envoy.
  * If update succeed there will be an ACK log for the instance.

## Run Manually

[How do I SET the GOPATH environment variable on Ubuntu? What file must I edit?](https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit)

[How To Install Go and Set Up a Local Programming Environment on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04)

[EnvironmentVariables](https://help.ubuntu.com/community/EnvironmentVariables)

to get docker host address we can use 

```
ip addr | grep eth0
```

run consul with docker:

``` bash
docker-compose -f docker-compose.consul.yaml up
```

``` bash
go get -u github.com/golang/dep/cmd/dep
go get -u github.com/derekparker/delve/cmd/dlv
mkdir ~/go/src/Envoy-Pilot
Gopkg.lock ~/go/src/Envoy-Pilot/
Gopkg.toml ~/go/src/Envoy-Pilot/
cd ~/go/src/Envoy-Pilot/
ls -l
dep ensure -vendor-only

cp -a cmd ~/go/src/Envoy-Pilot/cmd
rm -rf  ~/go/src/github.com/envoyproxy/go-control-plane/vendor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ~/go/bin/Envoy-Pilot ~/go/src/Envoy-Pilot/cmd/server/main.go

mkdir -p /go/bin/
bash ~/go/bin/Envoy-Pilot
```

Also we need to create a [.env](cmd/server/.env) file in root our app that use [godotenv package](https://github.com/joho/godotenv)

then we should run our apps container in this [docker-compose.apps.yaml](./docker-compose.apps.yaml) file on `8123`, `8126` ports.

Run Envoy Proxy with the following configurations or use `--service-node` && `--service-cluster`. they do same things.

[core.Node](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#core-node)

``` yaml
node:
  id: ride-service-replica-2
  cluster: ride-service
```

``` bash
envoy -c ./envoy-config.yaml -l info --service-cluster ride-service --service-node ride-service-replica-2

## or we can use node element in our envoy configs instead of --service-cluster, --service-node

envoy -c ./envoy-config.yaml -l info 
```


## Running Docker Compose

From root directory 
```
docker network create envoy-pilot_xds-demo
docker-compose -f docker-compose.consul.yaml up
docker-compose -f docker-compose.server.yaml up --build
```

## Running test

```
cd test/integration
docker-compose up --build

cd test/rspec
DEVMODE=true rspec basic_spec.rb --order defined --format documentation
```


## Runnnig Docker

Consul url need to be set in .env

env_values.txt
```
CONSUL_PATH="http://consul-server:8500"
CONSUL_PREFIX=""xDS"
```

Docker run
```
docker run -v $(pwd)/env_values.txt:/.env -p 7777:7777 -p 9090:9090 tak2siva/envoy-pilot:latest
```

## Helm Chart

Install using the [Helm Chart for Envoy-Pilot](https://github.com/tak2siva/Envoy-Pilot-Helm).

## Debugging

* xDS-Server is running on port 7777
* A http server is running on port 9090 for debugging

`localhost:9090/dump/KEY_TEMPLATE` will give a json dump of proto mapping

  **Ex:** 
  ```
  http://localhost:9090/dump/xDS/app-cluster/ride-service/(CDS|LDS|RDS|EDS)/config
  ```

* To get list of subcsribers hit `localhost:9090/dump/subscribers/`


## Prometheus metrics

Prometheus is running on `localhost:8081/metrics` and the following stats are available
  * xds_active_connections[cluster] (GAUGE)
  * xds_active_subscribers[cluster][type] (GAUGE)
  * xds_update_counter[cluster][subscribedTo] (Counter)
