package ctok

import (
	"net"

	corev1 "k8s.io/api/core/v1"

	"github.com/flomesh-io/fsm/pkg/connector"
)

var (
	syncCloudNamespace string
	withGateway        bool
)

// SetSyncCloudNamespace sets sync namespace
func SetSyncCloudNamespace(ns string) {
	syncCloudNamespace = ns
}

// IsSyncCloudNamespace if sync namespace
func IsSyncCloudNamespace(ns *corev1.Namespace) bool {
	if ns != nil {
		_, exists := ns.Annotations[connector.AnnotationMeshServiceSync]
		return exists
	}
	return false
}

// To4 converts the IPv4 address ip to a 4-byte representation.
// If ip is not an IPv4 address, To4 returns nil.
func (addr MicroEndpointAddr) To4() net.IP {
	return net.ParseIP(string(addr)).To4()
}

// To16 converts the IP address ip to a 16-byte representation.
// If ip is not an IP address (it is the wrong length), To16 returns nil.
func (addr MicroEndpointAddr) To16() net.IP {
	return net.ParseIP(string(addr)).To16()
}

// WithGateway sets enable or disable
func WithGateway(enable bool) {
	withGateway = enable
}
