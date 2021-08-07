module github.com/itzmanish/slatomate

go 1.16

// Temporary fix for etcd with grpc 1.27
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

require (
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/itzmanish/go-micro/v2 v2.10.0
	google.golang.org/protobuf v1.27.1
)
