module bbs-web

go 1.21.4

require (
	github.com/IBM/sarama v1.43.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/ecodeclub/ekit v0.0.8-0.20231107102407-45c43657fd3c
	github.com/gin-gonic/gin v1.9.1
	github.com/go-kratos/aegis v0.2.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.14.0
	github.com/go-sql-driver/mysql v1.7.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/mock v1.4.4
	github.com/google/wire v0.6.0
	github.com/gotomicro/redis-lock v0.0.3
	github.com/magiconair/properties v1.8.7
	github.com/mojocn/base64Captcha v1.3.6
	github.com/prometheus/client_golang v1.19.0
	github.com/redis/go-redis/v9 v9.5.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.9.0
	go.etcd.io/etcd/client/v3 v3.5.10
	go.mongodb.org/mongo-driver v1.14.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.49.0
	go.opentelemetry.io/otel v1.24.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.uber.org/mock v0.3.0
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.21.0
	golang.org/x/sync v0.6.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.32.0
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.8
	gorm.io/plugin/opentelemetry v0.1.4
	gorm.io/plugin/prometheus v0.1.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.7.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.etcd.io/etcd/api/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.10 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/image v0.13.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
