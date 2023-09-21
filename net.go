package compare

import (
	"fmt"
	"net"

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
	conn net.Conn
}

func (c *Connection) Write(msg string) error {
	if _, err := fmt.Fprint(c.conn, msg); err != nil {
		return err
	}
	return nil
}

func (t *Net) Open(addr string) (Connection, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return Connection{}, err
	}
	return Connection{
		conn: conn,
	}, nil
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.networker,
	}
}
