package driver

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type Driver struct {
	name     string
	region   string
	endpoint string

	server *grpc.Server

	ready bool

	csi.UnimplementedNodeServer
	csi.UnimplementedControllerServer
	csi.UnimplementedIdentityServer
}

type InputParams struct {
	Name     string
	Endpoint string
	Token    string
	Region   string
}

func NewDriver(params InputParams) *Driver {
	return &Driver{
		name:     params.Name,
		endpoint: params.Endpoint,
		region:   params.Endpoint,
	}
}

func (d *Driver) Run() error {
	url, err := url.Parse(d.endpoint)
	if err != nil {
		return fmt.Errorf("parsing the endpoint %s", err.Error())
	}

	if url.Scheme != "unix" {
		return fmt.Errorf("supported scheme is unix, provided %s", url.Scheme)
	}

	grpcAddress := path.Join(url.Host, filepath.FromSlash(url.Path))
	if url.Host == "" {
		grpcAddress = filepath.FromSlash(url.Path)
	}

	if err := os.Remove(grpcAddress); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing listen address %s", err.Error())
	}

	listener, err := net.Listen(url.Scheme, grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen %s", err.Error())
	}
	fmt.Print(listener)

	d.server = grpc.NewServer()
	csi.RegisterNodeServer(d.server, d)
	csi.RegisterControllerServer(d.server, d)
	csi.RegisterIdentityServer(d.server, d)

	d.ready = true

	return d.server.Serve(listener)
}
