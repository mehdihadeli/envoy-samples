This example is using Envoy-Pilot xDS Server (https://github.com/tak2siva/Envoy-Pilot) to load config from file

This sample don't use Consul and read envoy dynamic configuration from `FOLDER_PATH=/file_mode_config` in [env_values.txt](./env_values.txt) file. it use grpc connection and read its data from our specified location in our grpc service.


This docker compose setup contains:

    - Envoy (port: 10000)
    - App1 (port: 8123)

To start 
```
docker-compose up
```

Verify by 

```
curl http://localhost:10000
this is app --ONE--
```
