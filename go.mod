module github.com/owncloud/ocis-pkg/v2

go 1.13

require (
	github.com/ascarter/requestid v0.0.0-20170313220838-5b76ab3d4aee
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/cs3org/reva v0.1.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/justinas/alice v1.2.0
	github.com/micro/cli/v2 v2.1.1
	github.com/micro/go-micro/v2 v2.0.0
	github.com/micro/go-plugins/wrapper/trace/opencensus/v2 v2.0.1
	github.com/prometheus/client_golang v1.2.1
	github.com/restic/calens v0.2.0 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	go.opencensus.io v0.22.3
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
