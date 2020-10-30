Watches the given JSON or YAML file and updates Clusters
stored in the Turbine Labs API at startup and whenever the file changes.

The file can be specified as a flag or as the only argument (but not both).

The structure of the JSON and YAML formats is equivalent. Each contains 0 or
more clusters identified by name, each containing 0 or more instances. For
example, as YAML:

``` yaml
- cluster: c1
  instances:
  - host: h1
	port: 8000
	metadata:
	- key: stage
	  value: prod
```

Alternatively as JSON:
``` json
[
  {
	"cluster": "c1",
	"instances": [
	  {
		"host": "h1",
		"port": 8000,
		"metadata": [
		  { "key": "stage", "value": "prod" }
		]
	  }
	]
  }
]
```
Note that when updating the file, care should be taken to make the modification
atomic. In practice, this means writing the updated file to a temporary location and
then moving/renaming the file to the watched path. Alternatively, the watched path
may be a symbolic link that is replaced with a reference to the updated file.

[[rotor/xds] static cluster/listener support (fixes #5254) ](https://github.com/turbinelabs/rotor/commit/7acd7cd08a3512d0424a6b742033abb5d50f295c)

we can use `ip addr | grep eth0` to get host address for our docker container and inner our envoy configuration we can't use `host.docker.internal` because it just work in context docker like `dockerfile` and `docker-compose` but we can't use it inner our envoy configuration and we should use `host address (eth0 ip)` or `localhost` or `IPV6 localhost address ([::1])` in order to access other docker container such as access to `containersol/hello-world` container on expose port `8001`. we can test it after get host ip with `ip addr | grep eth0`.


docker run -v $(pwd)/:/data    \
  -e 'ROTOR_CMD=file' \
  -e 'ROTOR_CONSOLE_LEVEL=debug' \
  -e 'ROTOR_XDS_STATIC_RESOURCES_CONFLICT_BEHAVIOR=overwrite' \
  -e 'ROTOR_FILE_FORMAT=yaml' \
  -e 'ROTOR_FILE_FILENAME=/data/clusters.yaml' \
  -e 'ROTOR_XDS_STANDALONE_PORT=81' \
  -e 'ROTOR_XDS_STATIC_RESOURCES_FILENAME=/data/static_resources.yml' \
  -p 50000:50000 \
  turbinelabs/rotor:0.19.0


docker run \
  -e 'ENVOY_XDS_HOST=172.27.71.209' \
  -e 'ENVOY_XDS_PORT=50000' \
  -p 9999:9999 \
  -p 80:80 \
  turbinelabs/envoy-simple:0.19.0

  