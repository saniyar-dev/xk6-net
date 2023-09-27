package compare

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/net", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// comparator is the exported type
		networker *Net
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu:        vu,
		networker: &Net{vu: vu},
	}
}

type Net struct {
	vu modules.VU
}

type Connection struct {
	Conn net.Conn
}

func (c *Connection) Write(msg string) error {
	if _, err := c.Conn.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

type DialerConfig struct {
	KeepAlive int64
}

func (d *DialerConfig) ParseDialer() (net.Dialer, error) {
	return net.Dialer{
		KeepAlive: time.Duration(d.KeepAlive * 1000000000),
	}, nil
}

func (t *Net) Open(addr string, input map[string]interface{}) (Connection, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return Connection{}, err
	}
	var dConf DialerConfig
	err = json.Unmarshal(jsonData, &dConf)
	if err != nil {
		return Connection{}, err
	}

	d, err := dConf.ParseDialer()
	if err != nil {
		return Connection{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return Connection{}, err
	}
	return Connection{
		Conn: conn,
	}, nil
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.networker,
	}
}
