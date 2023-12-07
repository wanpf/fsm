package cache

import (
	"sync"

	gwpav1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policyattachment/v1alpha1"

	"github.com/flomesh-io/fsm/pkg/gateway/utils"
)

// UpstreamTLSPoliciesTrigger is responsible for processing TLSRoute objects
type UpstreamTLSPoliciesTrigger struct {
	mu sync.Mutex
}

// Insert adds a TLSRoute to the cache and returns true if the route is effective
func (p *UpstreamTLSPoliciesTrigger) Insert(obj interface{}, cache *GatewayCache) bool {
	policy, ok := obj.(*gwpav1alpha1.UpstreamTLSPolicy)
	if !ok {
		log.Error().Msgf("unexpected object type %T", obj)
		return false
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	cache.upstreamstls[utils.ObjectKey(policy)] = struct{}{}

	return cache.isRoutableTargetService(policy, policy.Spec.TargetRef)
}

// Delete removes a TLSRoute from the cache and returns true if the route was found
func (p *UpstreamTLSPoliciesTrigger) Delete(obj interface{}, cache *GatewayCache) bool {
	policy, ok := obj.(*gwpav1alpha1.UpstreamTLSPolicy)
	if !ok {
		log.Error().Msgf("unexpected object type %T", obj)
		return false
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	key := utils.ObjectKey(policy)
	_, found := cache.upstreamstls[key]
	delete(cache.upstreamstls, key)

	return found
}
