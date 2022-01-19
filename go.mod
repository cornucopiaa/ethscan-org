module eth2-exporter

go 1.18

require (
	cloud.google.com/go v0.72.0
	firebase.google.com/go v3.13.0+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/chromedp/cdproto v0.0.0-20200709115526-d1f6fc58448b
	github.com/chromedp/chromedp v0.5.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ethereum/go-ethereum v1.9.14
	github.com/evanw/esbuild v0.8.23
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/context v1.1.1
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.2.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jackc/pgx/v4 v4.6.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/juliangruber/go-intersect v1.0.1-0.20200323101606-4bd944a17692
	github.com/kataras/i18n v0.0.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.2.0
	github.com/mailgun/mailgun-go/v4 v4.1.3
	github.com/mssola/user_agent v0.5.2
	github.com/phyber/negroni-gzip v0.0.0-20180113114010-ef6356a5d029
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/protolambda/zrnt v0.12.4
	github.com/protolambda/ztyp v0.1.0
	github.com/prysmaticlabs/eth2-types v0.0.0-20210303084904-c9735a06829d
	github.com/prysmaticlabs/ethereumapis v0.0.0-20200909190233-a9190508298e
	github.com/prysmaticlabs/go-bitfield v0.0.0-20200618145306-2ae0807bef65
	github.com/prysmaticlabs/prysm v1.0.0-alpha.25.0.20200917185001-3db678499074
	github.com/sirupsen/logrus v1.6.0
	github.com/stripe/stripe-go/v72 v72.24.0
	github.com/swaggo/http-swagger v1.3.3
	github.com/swaggo/swag v1.8.9
	github.com/urfave/negroni v1.0.0
	github.com/zesik/proxyaddr v0.0.0-20161218060608-ec32c535184d
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/text v0.3.7
	google.golang.org/api v0.36.0
	google.golang.org/genproto v0.0.0-20201203001206-6486ece9c497
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cloud.google.com/go/firestore v1.4.0 // indirect
	cloud.google.com/go/storage v1.10.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/ferranbt/fastssz v0.0.0-20210120143747-11b9eff30ea9 // indirect
	github.com/go-chi/chi v4.0.2+incompatible // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.2 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/herumi/bls-eth-go-binary v0.0.0-20200722032157-41fc56eba7b4 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.5.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200307190119-3430c5407db8 // indirect
	github.com/jackc/pgtype v1.3.0 // indirect
	github.com/jackc/puddle v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/knq/sysutil v0.0.0-20191005231841-15668db23d08 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/minio/sha256-simd v0.1.1 // indirect
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	go.opencensus.io v0.22.5 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/ethereum/go-ethereum => github.com/prysmaticlabs/bazel-go-ethereum v0.0.0-20200626171358-a933315235ec
