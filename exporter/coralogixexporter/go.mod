module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter

go 1.18

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.69.0
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.69.2-0.20230112233839-f2a0133bf677
	go.opentelemetry.io/collector/component v0.69.2-0.20230112233839-f2a0133bf677
	go.opentelemetry.io/collector/confmap v0.69.2-0.20230112233839-f2a0133bf677
	go.opentelemetry.io/collector/consumer v0.69.2-0.20230112233839-f2a0133bf677
	go.opentelemetry.io/collector/pdata v1.0.0-rc3.0.20230112233839-f2a0133bf677
	google.golang.org/genproto v0.0.0-20221202195650-67e5cbc046fd
	google.golang.org/grpc v1.52.0
)

require (
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.14 // indirect
	github.com/knadh/koanf v1.4.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.1.17 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.69.2-0.20230112233839-f2a0133bf677 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.37.0 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/metric v0.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger => ../../pkg/translator/jaeger

retract v0.65.0
