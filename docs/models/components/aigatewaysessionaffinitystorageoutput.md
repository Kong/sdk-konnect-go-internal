# AIGatewaySessionAffinityStorageOutput

Where session affinity state is stored.


## Supported Types

### AIGatewaySessionAffinityStorageLocal

```go
aiGatewaySessionAffinityStorageOutput := components.CreateAIGatewaySessionAffinityStorageOutputLocal(components.AIGatewaySessionAffinityStorageLocal{/* values here */})
```

### AIGatewaySessionAffinityStorageRedisOutput

```go
aiGatewaySessionAffinityStorageOutput := components.CreateAIGatewaySessionAffinityStorageOutputRedis(components.AIGatewaySessionAffinityStorageRedisOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewaySessionAffinityStorageOutput.Type {
	case components.AIGatewaySessionAffinityStorageOutputTypeLocal:
		// aiGatewaySessionAffinityStorageOutput.AIGatewaySessionAffinityStorageLocal is populated
	case components.AIGatewaySessionAffinityStorageOutputTypeRedis:
		// aiGatewaySessionAffinityStorageOutput.AIGatewaySessionAffinityStorageRedisOutput is populated
}
```
