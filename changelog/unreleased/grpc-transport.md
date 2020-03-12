Bugfix: Tracing middleware

When creating a grpc service the default transport being used was http, instead of gRPC.

https://github.com/owncloud/ocis-pkg/issues/38
https://github.com/owncloud/ocis-pkg/pull/39
