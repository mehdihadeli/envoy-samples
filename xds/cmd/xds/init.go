package main

import (
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/dubbo/router/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/adaptive_concurrency/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/aws_lambda/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/aws_request_signing/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/buffer/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/cache/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/compressor/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/cors/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/csrf/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/dynamic_forward_proxy/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/dynamo/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/ext_authz/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/fault/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/grpc_http1_bridge/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/grpc_stats/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/grpc_web/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/gzip/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/header_to_metadata/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/health_check/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/ip_tagging/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/jwt_authn/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/lua/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/on_demand/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/original_src/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/rate_limit/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/rbac/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/router/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/squash/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/tap/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/transcoder/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/http_inspector/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/original_dst/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/original_src/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/proxy_protocol/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/tls_inspector/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/client_ssl_auth/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/direct_response/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/dubbo_proxy/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/echo/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/ext_authz/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/kafka_broker/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/local_rate_limit/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/mongo_proxy/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/mysql_proxy/v1alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/rate_limit/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/rbac/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/redis_proxy/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/sni_cluster/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/thrift_proxy/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/zookeeper_proxy/v1alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/thrift/rate_limit/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/thrift/router/v2alpha1"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/filter/udp/udp_proxy/v2alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/fault/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/adaptive_concurrency/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/admission_control/v3alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_lambda/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_request_signing/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/buffer/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cache/v3alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cors/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/csrf/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/decompressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamo/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_authz/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_bridge/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_reverse_bridge/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_json_transcoder/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_stats/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_web/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/gzip/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/header_to_metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/health_check/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ip_tagging/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/lua/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/on_demand/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/original_src/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/rbac/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/squash/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/http_inspector/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_dst/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_src/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/proxy_protocol/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/client_ssl_auth/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/direct_response/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/router/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/echo/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ext_authz/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/kafka_broker/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/local_ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/mongo_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/mysql_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/postgres_proxy/v3alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/rbac/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/redis_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/rocketmq_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_cluster/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_dynamic_forward_proxy/v3alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/zookeeper_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/dns_filter/v3alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/udp_proxy/v3"
)
