# Envoy-Pilot 

[![Build Status](https://travis-ci.org/tak2siva/Envoy-Pilot.svg?branch=master)](https://travis-ci.org/tak2siva/Envoy-Pilot)    ![Version](https://img.shields.io/badge/version-v0.2.3-yellowgreen.svg)
![docker pulls](https://img.shields.io/docker/pulls/tak2siva/envoy-pilot)


Envoy Pilot or Envoy xDS is a control plane implementation for [Envoy](https://github.com/envoyproxy/envoy) written in Golang and uses Consul for persistence by default. It can also run without Consul by loading configuration from file.

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

Run Envoy Proxy with the following configurations or use `--service-node` && `--service-cluster`
```
node:
  id: ride-service-replica-2
  cluster: ride-service
```

Every *DS requires two keys to be set in consul
  * config
  * version

And the key template is `xDS/app-cluster/ride-service/CDS/(config|version)`

For CDS add KV pairs
  * `xDS/app-cluster/ride-service/CDS/version` => "1.0"
  * `xDS/app-cluster/ride-service/CDS/config` => `"[{
      {
        "name": "app1",
        "connect_timeout": "0.250s",
        "type": "STRICT_DNS",
        "lb_policy": "RANDOM",
        "hosts": [{
          "socket_address": {
           "address": "127.0.0.2",
           "port_value": 1234
          }
        }]
    }
  }]"`

Pushing new configuration
  * Envoy-Pilot will be polling for version change every 10 seconds.  
  * If there is a version mismatch for any of `xDS/app-cluster/ride-service/(CDS|LDS|RDS|EDS)/version` then new config `xDS/app-cluster/ride-service/(CDS|LDS|RDS|EDS)/config` will be pushed to subscriber envoy.
  * If update succeed there will be an ACK log for the instance.

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
