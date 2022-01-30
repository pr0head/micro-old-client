module github.com/pr0head/micro-old-client

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/client/selector/static v0.0.0-20200119172437-4fe21aa238fd
	github.com/pr0head/micro-old-server v0.0.0-20220130073554-1eab463906de
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
