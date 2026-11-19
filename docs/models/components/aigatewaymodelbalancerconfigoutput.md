# AIGatewayModelBalancerConfigOutput

Configuration for a model's load balancer when multiple target models are configured.


## Supported Types

### AIGatewayModelBalancerConsistentHashingConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputConsistentHashing(components.AIGatewayModelBalancerConsistentHashingConfig{/* values here */})
```

### AIGatewayModelBalancerIntelligentConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputIntelligent(components.AIGatewayModelBalancerIntelligentConfigOutput{/* values here */})
```

### AIGatewayModelBalancerLeastConnectionsConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLeastConnections(components.AIGatewayModelBalancerLeastConnectionsConfigOutput{/* values here */})
```

### AIGatewayModelBalancerLowestLatencyConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLowestLatency(components.AIGatewayModelBalancerLowestLatencyConfigOutput{/* values here */})
```

### AIGatewayModelBalancerLowestUsageConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLowestUsage(components.AIGatewayModelBalancerLowestUsageConfigOutput{/* values here */})
```

### AIGatewayModelBalancerPriorityConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputPriority(components.AIGatewayModelBalancerPriorityConfigOutput{/* values here */})
```

### AIGatewayModelBalancerRoundRobinConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputRoundRobin(components.AIGatewayModelBalancerRoundRobinConfigOutput{/* values here */})
```

### AIGatewayModelBalancerSemanticConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputSemantic(components.AIGatewayModelBalancerSemanticConfigOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerConfigOutput.Type {
	case components.AIGatewayModelBalancerConfigOutputTypeConsistentHashing:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerConsistentHashingConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeIntelligent:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerIntelligentConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLeastConnections:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLeastConnectionsConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLowestLatency:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLowestLatencyConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLowestUsage:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLowestUsageConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypePriority:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerPriorityConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypeRoundRobin:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerRoundRobinConfigOutput is populated
	case components.AIGatewayModelBalancerConfigOutputTypeSemantic:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerSemanticConfigOutput is populated
}
```
