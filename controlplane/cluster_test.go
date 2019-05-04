package controlplane_test

import (
	"testing"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	cp "github.com/gojekfarm/envoy-lb-operator/controlplane"
	"github.com/stretchr/testify/assert"
)

func TestStrictDNSLRCluster(t *testing.T) {
	cluster := cp.StrictDNSLRCluster("test", "svc", uint32(443), 1000)
	assert.Equal(t, "test", cluster.Name)
	assert.Equal(t, v2.Cluster_STRICT_DNS, cluster.ClusterDiscoveryType.(*v2.Cluster_Type).Type)
	assert.Equal(t, v2.Cluster_LEAST_REQUEST, cluster.LbPolicy)
	assert.Equal(t, "test", cluster.LoadAssignment.ClusterName)
	assert.Equal(t, 1, len(cluster.LoadAssignment.Endpoints))
	assert.Equal(t, 1, len(cluster.LoadAssignment.Endpoints[0].LbEndpoints))
	lbe := cluster.LoadAssignment.Endpoints[0].LbEndpoints[0]
	ep := lbe.HostIdentifier.(*endpoint.LbEndpoint_Endpoint).Endpoint
	socketAddress := ep.Address.Address.(*core.Address_SocketAddress).SocketAddress
	assert.Equal(t, "svc", socketAddress.Address)
	assert.Equal(t, uint32(443), socketAddress.PortSpecifier.(*core.SocketAddress_PortValue).PortValue)

}

func TestStrictDNSLRHttp2Cluster(t *testing.T) {
	cluster := cp.StrictDNSLRHttp2Cluster("test", "svc", uint32(443), 1000)
	assert.Equal(t, "test", cluster.Name)
	assert.Equal(t, v2.Cluster_STRICT_DNS, cluster.ClusterDiscoveryType.(*v2.Cluster_Type).Type)
	assert.Equal(t, v2.Cluster_LEAST_REQUEST, cluster.LbPolicy)
	assert.NotNil(t, cluster.Http2ProtocolOptions)
	assert.Equal(t, "test", cluster.LoadAssignment.ClusterName)
	assert.Equal(t, 1, len(cluster.LoadAssignment.Endpoints))
	assert.Equal(t, 1, len(cluster.LoadAssignment.Endpoints[0].LbEndpoints))
	lbe := cluster.LoadAssignment.Endpoints[0].LbEndpoints[0]
	ep := lbe.HostIdentifier.(*endpoint.LbEndpoint_Endpoint).Endpoint
	socketAddress := ep.Address.Address.(*core.Address_SocketAddress).SocketAddress
	assert.Equal(t, "svc", socketAddress.Address)
	assert.Equal(t, uint32(443), socketAddress.PortSpecifier.(*core.SocketAddress_PortValue).PortValue)

}
