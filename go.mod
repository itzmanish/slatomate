module github.com/itzmanish/slatomate

go 1.16

// Temporary fix for etcd with grpc 1.27
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/briandowns/spinner v1.16.0
	github.com/fatih/color v1.12.0
	github.com/google/uuid v1.3.0
	github.com/graphql-go/graphql v0.8.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/itzmanish/go-micro-plugins/wrapper/trace/opentracing/v2 v2.10.0
	github.com/itzmanish/go-micro/v2 v2.10.1
	github.com/joho/godotenv v1.3.0
	github.com/manifoldco/promptui v0.8.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/olekukonko/tablewriter v0.0.4
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.0
	github.com/slack-go/slack v0.9.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/ysugimoto/grpc-graphql-gateway v0.20.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.12
)
