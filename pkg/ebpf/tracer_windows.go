// +build windows

package ebpf

import (
	"sync"
)

// CurrentKernelVersion is not implemented on non-linux systems
func CurrentKernelVersion() (uint32, error) {
	return 1, nil
}

type Tracer struct {
	config *Config
	portMapping *PortMapping
	// Telemetry
	perfReceived  int64
	perfLost      int64
	skippedConns  int64
	pidCollisions int64
	// Will track the count of expired TCP connections
	// We are manually expiring TCP connections because it seems that we are losing some of the TCP close events
	// For now we are only tracking the `tcp_close` probe but we should also track the `tcp_set_state` probe when
	// states are set to `TCP_CLOSE_WAIT`, `TCP_CLOSE` and `TCP_FIN_WAIT1` we should probably also track `tcp_time_wait`
	// However there are some caveats by doing that:
	// - `tcp_set_state` does not have access to the PID of the connection => we have to remove the PID from the connection tuple (which can lead to issues)
	// - We will have multiple probes that can trigger for the same connection close event => we would have to add something to dedupe those
	// - - Using the timestamp does not seem to be reliable (we are already seeing unordered connections)
	// - - Having IDs for those events would need to have an internal monotonic counter and this is tricky to manage (race conditions, cleaning)
	//
	// If we want to have a way to track the # of active TCP connections in the future we could use the procfs like here: https://github.com/DataDog/datadog-agent/pull/3728
	// to determine whether a connection is truly closed or not
	expiredTCPConns int64
	buffer     []ConnectionStats
	bufferLock sync.Mutex
	// Connections for the tracer to blacklist
	sourceExcludes []*ConnectionFilter
	destExcludes   []*ConnectionFilter

}

func NewTracer(config *Config) (*Tracer, error) {
	return &Tracer{}, nil
}

func (t *Tracer) Stop() {}

func (t *Tracer) GetActiveConnections(_ string) (*Connections, error) {
	return nil, ErrNotImplemented
}

func (t *Tracer) GetStats() (map[string]interface{}, error) {
	return nil, ErrNotImplemented
}

func (t *Tracer) DebugNetworkState(clientID string) (map[string]interface{}, error) {
	return nil, ErrNotImplemented
}

func (t *Tracer) DebugNetworkMaps() (*Connections, error) {
	return nil, ErrNotImplemented
}

