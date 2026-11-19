# AIGatewaySessionAffinityStorage

Where session affinity state is stored.


## Supported Types

### AIGatewaySessionAffinityStorageLocal

```go
aiGatewaySessionAffinityStorage := components.CreateAIGatewaySessionAffinityStorageLocal(components.AIGatewaySessionAffinityStorageLocal{/* values here */})
```

### AIGatewaySessionAffinityStorageRedis

```go
aiGatewaySessionAffinityStorage := components.CreateAIGatewaySessionAffinityStorageRedis(components.AIGatewaySessionAffinityStorageRedis{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewaySessionAffinityStorage.Type {
	case components.AIGatewaySessionAffinityStorageTypeLocal:
		// aiGatewaySessionAffinityStorage.AIGatewaySessionAffinityStorageLocal is populated
	case components.AIGatewaySessionAffinityStorageTypeRedis:
		// aiGatewaySessionAffinityStorage.AIGatewaySessionAffinityStorageRedis is populated
}
```
