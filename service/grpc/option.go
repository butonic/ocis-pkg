package grpc

import (
	"context"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/transport"
	grpcT "github.com/micro/go-micro/v2/transport/grpc"
	"github.com/owncloud/ocis-pkg/v2/log"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Logger    log.Logger
	Namespace string
	Name      string
	Version   string
	Address   string
	Context   context.Context
	Client    client.Client
	Transport transport.Transport
	Flags     []cli.Flag
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{
		Namespace: "go.micro.api",
	}

	for _, o := range opts {
		o(&opt)
	}

	if opt.Client == nil {
		opt.Client = grpc.NewClient()
	}

	if opt.Transport == nil {
		opt.Transport = grpcT.NewTransport()
	}

	return opt
}

// Logger provides a function to set the logger option.
func Logger(l log.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// Namespace provides a function to set the namespace option.
func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

// Name provides a function to set the name option.
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Version provides a function to set the version option.
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

// Address provides a function to set the address option.
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Context provides a function to set the context option.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// Flags provides a function to set the flags option.
func Flags(flags ...cli.Flag) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// Client provides a function to set the client option.
func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

// Transport provides a function to set the transport option.
func Transport(t transport.Transport) Option {
	return func(o *Options) {
		o.Transport = t
	}
}
