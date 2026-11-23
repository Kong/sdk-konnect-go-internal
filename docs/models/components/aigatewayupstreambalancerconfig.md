# AIGatewayUpstreamBalancerConfig

Load balancing configuration used to distribute requests across the upstream targets.


## Supported Types

### AIGatewayUpstreamBalancerConsistentHashingConfig

```go
aiGatewayUpstreamBalancerConfig := components.CreateAIGatewayUpstreamBalancerConfigConsistentHashing(components.AIGatewayUpstreamBalancerConsistentHashingConfig{/* values here */})
```

### AIGatewayUpstreamBalancerLeastConnectionsConfig

```go
aiGatewayUpstreamBalancerConfig := components.CreateAIGatewayUpstreamBalancerConfigLeastConnections(components.AIGatewayUpstreamBalancerLeastConnectionsConfig{/* values here */})
```

### AIGatewayUpstreamBalancerLowestLatencyConfig

```go
aiGatewayUpstreamBalancerConfig := components.CreateAIGatewayUpstreamBalancerConfigLowestLatency(components.AIGatewayUpstreamBalancerLowestLatencyConfig{/* values here */})
```

### AIGatewayUpstreamBalancerRoundRobinConfig

```go
aiGatewayUpstreamBalancerConfig := components.CreateAIGatewayUpstreamBalancerConfigRoundRobin(components.AIGatewayUpstreamBalancerRoundRobinConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayUpstreamBalancerConfig.Type {
	case components.AIGatewayUpstreamBalancerConfigTypeConsistentHashing:
		// aiGatewayUpstreamBalancerConfig.AIGatewayUpstreamBalancerConsistentHashingConfig is populated
	case components.AIGatewayUpstreamBalancerConfigTypeLeastConnections:
		// aiGatewayUpstreamBalancerConfig.AIGatewayUpstreamBalancerLeastConnectionsConfig is populated
	case components.AIGatewayUpstreamBalancerConfigTypeLowestLatency:
		// aiGatewayUpstreamBalancerConfig.AIGatewayUpstreamBalancerLowestLatencyConfig is populated
	case components.AIGatewayUpstreamBalancerConfigTypeRoundRobin:
		// aiGatewayUpstreamBalancerConfig.AIGatewayUpstreamBalancerRoundRobinConfig is populated
}
```
