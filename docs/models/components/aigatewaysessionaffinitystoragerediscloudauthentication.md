# AIGatewaySessionAffinityStorageRedisCloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthenticationOutput

```go
aiGatewaySessionAffinityStorageRedisCloudAuthentication := components.CreateAIGatewaySessionAffinityStorageRedisCloudAuthenticationAws(components.AIGatewayRedisAWSAuthenticationOutput{/* values here */})
```

### AIGatewayRedisAzureAuthenticationOutput

```go
aiGatewaySessionAffinityStorageRedisCloudAuthentication := components.CreateAIGatewaySessionAffinityStorageRedisCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthenticationOutput{/* values here */})
```

### AIGatewayRedisGCPAuthenticationOutput

```go
aiGatewaySessionAffinityStorageRedisCloudAuthentication := components.CreateAIGatewaySessionAffinityStorageRedisCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthenticationOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewaySessionAffinityStorageRedisCloudAuthentication.Type {
	case components.AIGatewaySessionAffinityStorageRedisCloudAuthenticationTypeAws:
		// aiGatewaySessionAffinityStorageRedisCloudAuthentication.AIGatewayRedisAWSAuthenticationOutput is populated
	case components.AIGatewaySessionAffinityStorageRedisCloudAuthenticationTypeAzure:
		// aiGatewaySessionAffinityStorageRedisCloudAuthentication.AIGatewayRedisAzureAuthenticationOutput is populated
	case components.AIGatewaySessionAffinityStorageRedisCloudAuthenticationTypeGcp:
		// aiGatewaySessionAffinityStorageRedisCloudAuthentication.AIGatewayRedisGCPAuthenticationOutput is populated
}
```
