module github.com/itzmanish/slatomate

go 1.16

// Temporary fix for etcd with grpc 1.27
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

require (
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/itzmanish/go-micro/v2 v2.10.0
	github.com/joho/godotenv v1.3.0
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/slack-go/slack v0.9.4
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.12
)
